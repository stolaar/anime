package scrapper

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"anime/models"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

var (
	opts = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU)
	BaseUrl = "https://anix.to"
)

func GetNumberOfEpisodes(animePath string) (string, error) {
	ctx, cancel := createChromedpContext()
	defer cancel()

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.ActionFunc(AddStealthScript),
			chromedp.Navigate(BaseUrl + animePath),
			chromedp.WaitReady("div#ani-episode span"),
			chromedp.Click("div#ani-episode span"),
			chromedp.Click("div#ani-episode span"),
			chromedp.Nodes("div#ani-episode a:last-child", &nodes),
		})
	if err != nil {
		return "", err
	}

	var result string

	if len(nodes) > 0 {
		node := nodes[0]
		if len(node.Children) > 0 {
			result = strings.Split(node.Children[0].NodeValue, "-")[1]
		}
	}

	return result, nil
}

func ScrapeAnimesByQuery(query string) ([]*models.Anime, error) {
	ctx, cancel := createChromedpContext()
	defer cancel()

	inputSelector := `input[name="keyword"]`

	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.ActionFunc(AddStealthScript),
			chromedp.Navigate(BaseUrl),
			chromedp.WaitReady(inputSelector),
			chromedp.Click(inputSelector),
			chromedp.Focus(inputSelector),
			chromedp.SendKeys(inputSelector, query),
			chromedp.WaitVisible(`a.piece`, chromedp.ByQueryAll),
			chromedp.Nodes(`a.piece`, &nodes),
		})
	if err != nil {
		return nil, err
	}
	result := []*models.Anime{}
	var getAnimesErr error

	for _, i := range nodes {
		var images []*cdp.Node
		var titles []*cdp.Node

		getImgs := chromedp.Nodes(`a.piece img`, &images, chromedp.ByQuery, chromedp.FromNode(i))
		getTitles := chromedp.Nodes(`.ani-name`, &titles, chromedp.ByQuery, chromedp.FromNode(i))
		if err := chromedp.Run(ctx, getImgs, getTitles); err != nil {
			getAnimesErr = err
			break
		}

		for idx, n := range images {
			result = append(result, &models.Anime{
				Image: n.AttributeValue("src"),
				Title: titles[idx].Children[0].NodeValue,
				Link:  i.AttributeValue("href"),
			})
		}
	}

	if getAnimesErr != nil {
		return nil, getAnimesErr
	}

	return result, nil
}

func ScrapeEpisode(url string) (*models.Episode, error) {
	ctx, cancel := createChromedpContext()
	defer cancel()

	var iframes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.ActionFunc(AddStealthScript),
			chromedp.Navigate(url),
			chromedp.WaitReady(".wrapper > main"),
			chromedp.Click("div#player > div"),
			chromedp.Click("div#player > div"),
			chromedp.Sleep(1 * time.Second),
			chromedp.Nodes(`iframe`, &iframes),
		})
	if err != nil {
		return nil, err
	}

	var result *models.Episode

	for _, i := range iframes {
		src := i.AttributeValue("src")
		if validSrc(src) {
			result = &models.Episode{
				Src: src,
			}
			break
		}
	}

	return result, nil
}

func validSrc(src string) bool {
	return strings.Contains(src, "vid")
}

func listenForNetworkEvent(ctx context.Context) map[string]string {
	urlMap := map[string]string{}
	var wg sync.WaitGroup

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventLoadingFinished:
			wg.Add(1)
			go func() {
				c := chromedp.FromContext(ctx)
				_, err := network.GetResponseBody(ev.RequestID).Do(cdp.WithExecutor(ctx, c.Target))
				if err != nil {
					defer wg.Done()
					return
				}

				defer wg.Done()
			}()

		case *network.EventResponseReceived:
			url := ev.Response.URL
			urlMap[ev.RequestID.String()] = url
		}
	})

	return urlMap
}

func createChromedpContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx, _ = chromedp.NewExecAllocator(ctx, opts...)
	ctx, _ = chromedp.NewContext(ctx,
		chromedp.WithLogf(log.Printf))

	return ctx, cancel
}

func AddStealthScript(ctx context.Context) error {
	var err error
	_, err = page.AddScriptToEvaluateOnNewDocument(scriptx).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
