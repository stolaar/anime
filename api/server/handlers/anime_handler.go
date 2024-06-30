package handlers

import (
	"net/http"
	"strconv"

	"anime/models"
	"anime/repositories"
	"anime/responses"
	s "anime/server"
	"anime/services/scrapper"

	"github.com/labstack/echo/v4"
)

type AnimeHandlers struct {
	server *s.Server
}

func NewAnimeHandlers(server *s.Server) *AnimeHandlers {
	return &AnimeHandlers{server: server}
}

// GetAnimes godoc
//
//	@Summary	 Get all animes
//	@Description Get animes
//	@ID				animes-create
//	@Tags			Animes Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.CreatePostRequest	true	"Post title and content"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/posts [post]
func (a *AnimeHandlers) GetAnimes(c echo.Context) error {
	query := c.QueryParam("q")
	var animes []models.Anime

	animeRepository := repositories.NewAnimeRepository(a.server.DB)
	animeRepository.GetAllAnimes(&animes)

	if len(animes) <= 0 {
		scrappedAnimes, _ := scrapper.ScrapeAnimesByQuery(query)

		for _, scrapped := range scrappedAnimes {
			animes = append(animes, models.Anime(*scrapped))
		}

		animeRepository.DB.Save(scrappedAnimes)

		animeRepository.GetAllAnimes(&animes)
	}

	response := responses.NewAnimeResponse(animes)
	return responses.Response(c, http.StatusOK, response)
}

// GetAnimeEpisodes godoc
//
//	@Summary		Get anime episodes
//	@Description	Get the list of all anime episodes
//	@ID				anime-episodes-get
//	@Tags			Get Actions
//	@Produce		json
//	@Success		200	{array}	responses.GetAnimes
//	@Security		ApiKeyAuth
//	@Router			/animes [get]
func (a *AnimeHandlers) GetAnimeEpisodes(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var episodes []models.Episode

	animeRepository := repositories.NewAnimeRepository(a.server.DB)
	animeRepository.GetEpisodes(id, &episodes)

	if len(episodes) <= 0 {
		var anime models.Anime
		animeRepository.DB.Model(&models.Anime{}).Select("link").Where("id = ?", id).First(&anime)

		lastEpisodeNumber, err := scrapper.GetNumberOfEpisodes(anime.Link)
		if err != nil {
			return err
		}
		num, err := strconv.Atoi(lastEpisodeNumber)
		if err != nil {
			return err
		}

		for i := 1; i <= num; i++ {
			si := strconv.Itoa(i)
			episodes = append(episodes, models.Episode{
				Title:   si,
				AnimeID: uint(id),
				Link:    scrapper.BaseUrl + anime.Link + "/" + "ep-" + si,
			})
		}

		animeRepository.DB.Save(episodes)
	}

	response := responses.NewEpisodesResponse(episodes)
	return responses.Response(c, http.StatusOK, response)
}

// GetAnimeEpisodeSrc godoc
//
//	@Summary		Get anime episode src
//	@Description	Get the link for watching the episode
//	@ID				anime-episodes-get
//	@Tags			Get Actions
//	@Produce		json
//	@Success		200	{array}	responses.GetAnimes
//	@Security		ApiKeyAuth
//	@Router			/animes [get]
func (a *AnimeHandlers) GetAnimeEpisodeSrc(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	animeRepository := repositories.NewAnimeRepository(a.server.DB)

	var episode models.Episode
	animeRepository.DB.Model(&models.Episode{}).Select("link").Where("id = ?", id).First(&episode)

	src, err := scrapper.ScrapeEpisode(episode.Link)
	if err != nil {
		return err
	}

	return responses.Response(c, http.StatusOK, src.Src)
}
