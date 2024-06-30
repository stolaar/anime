package anime

import "anime/models"

func (animeService *Service) Create(anime *models.Anime) {
	animeService.DB.Create(anime)
}
