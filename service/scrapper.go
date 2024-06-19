package service

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/stolaar/anime/model"
)

var (
	opts = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU)
	baseUrl = "https://anix.to"
)

func GetNumberOfEpisodes(animePath string) (string, error) {
	ctx, cancel := createChromedpContext()
	defer cancel()

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.ActionFunc(AddStealthScript),
			chromedp.Navigate(baseUrl + animePath),
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

func ScrapeAnimesByQuery(query string) ([]*model.Anime, error) {
	ctx, cancel := createChromedpContext()
	defer cancel()

	inputSelector := `input[name="keyword"]`

	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.ActionFunc(AddStealthScript),
			chromedp.Navigate(baseUrl),
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
	result := []*model.Anime{}
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
			href := i.AttributeValue("href")
			result = append(result, &model.Anime{
				Id:    strings.Replace(strings.ReplaceAll(href, "/", "-"), "-", "", 1),
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

func ScrapeEpisode(url string) (*model.Episode, error) {
	ctx, cancel := createChromedpContext()
	defer cancel()

	var iframes []*cdp.Node
	listenForNetworkEvent(ctx)
	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.ActionFunc(AddStealthScript),
			chromedp.Navigate(url),
			chromedp.WaitReady(".content"),
			chromedp.Click(".player-btn"),
			chromedp.Click(".player-btn"),
			chromedp.Sleep(1 * time.Second),
			chromedp.Nodes(`iframe`, &iframes),
		})
	if err != nil {
		return nil, err
	}

	var result *model.Episode

	for _, i := range iframes {
		src := i.AttributeValue("src")
		if validSrc(src) {
			result = &model.Episode{
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
			println("Response URL", url)
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

const scriptx = `(function(w, n, wn) {
	// Pass the Webdriver Test.
	Object.defineProperty(n, 'webdriver', {
	  get: () => false,
	});

	// Pass the Plugins Length Test.
	// Overwrite the plugins property to use a custom getter.
	Object.defineProperty(n, 'plugins', {
	  // This just needs to have length > 0 for the current test,
	  // but we could mock the plugins too if necessary.
	  get: () => [1, 2, 3, 4, 5],
	});

	// Pass the Languages Test.
	// Overwrite the plugins property to use a custom getter.
	Object.defineProperty(n, 'languages', {
	  get: () => ['en-US', 'en'],
	});

	// Pass the Chrome Test.
	// We can mock this in as much depth as we need for the test.
	w.chrome = {
	  runtime: {},
	};

	// Pass the Permissions Test.
	const originalQuery = wn.permissions.query;
	return wn.permissions.query = (parameters) => (
	  parameters.name === 'notifications' ?
		Promise.resolve({ state: Notification.permission }) :
		originalQuery(parameters)
	);

  })(window, navigator, window.navigator);`
