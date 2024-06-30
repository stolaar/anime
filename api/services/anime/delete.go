package anime

import "anime/models"

func (animeService *Service) Delete(anime *models.Anime) {
	animeService.DB.Delete(anime)
}
