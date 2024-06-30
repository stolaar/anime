package anime

import (
	"anime/requests"

	"anime/models"

	"gorm.io/gorm"
)

type ServiceWrapper interface {
	Create(anime *models.Anime)
	Delete(anime *models.Anime)
	Update(anime *models.Anime, updateAnimeRequest *requests.UpdatePostRequest)
}

type Service struct {
	DB *gorm.DB
}

func (animeService *Service) FindAnimes() {
}

func NewPostService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
