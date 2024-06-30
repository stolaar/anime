package anime

import (
	"anime/requests"

	"anime/models"
)

func (animeService *Service) Update(anime *models.Anime, updateAnimeRequest *requests.UpdateAnimeRequest) {
	anime.Title = updateAnimeRequest.Title
	anime.Src = updateAnimeRequest.Src
	animeService.DB.Save(anime)
}
