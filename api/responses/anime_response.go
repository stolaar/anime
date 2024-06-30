package responses

import (
	"anime/models"
)

type AnimeResponse struct {
	Title       string `json:"title" example:"Echo"`
	Description string `json:"description" example:"Echo is nice!"`
	Image       string `json:"image" example:"John Doe"`
	Link        string `json:"link" example:"John Doe"`
	ID          uint   `json:"id" example:"1"`
	Src         string `json:"src" example:"1"`
}

type EpisodesResponse struct {
	models.Episode
	ID uint `json:"id" example:"1"`
}

func NewAnimeResponse(animes []models.Anime) *[]AnimeResponse {
	postResponse := make([]AnimeResponse, 0)

	for i := range animes {
		postResponse = append(postResponse, AnimeResponse{
			Title:       animes[i].Title,
			Description: animes[i].Description,
			Image:       animes[i].Image,
			Link:        animes[i].Link,
			Src:         animes[i].Src,
			ID:          animes[i].ID,
		})
	}

	return &postResponse
}

func NewEpisodesResponse(episodes []models.Episode) *[]EpisodesResponse {
	response := make([]EpisodesResponse, 0)

	for _, episode := range episodes {
		ID := episode.ID
		ep := EpisodesResponse{
			episode,
			ID,
		}
		response = append(response, ep)
	}
	return &response
}
