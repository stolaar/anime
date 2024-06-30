package repositories

import (
	"anime/models"

	"gorm.io/gorm"
)

type AnimeRepositoryQ interface {
	GetPosts(posts *[]models.Post)
	GetPost(post *models.Post, id int)
}

type AnimeRepository struct {
	DB *gorm.DB
}

func NewAnimeRepository(db *gorm.DB) *AnimeRepository {
	return &AnimeRepository{DB: db}
}

func (animeRepository *AnimeRepository) GetEpisodes(id int, episodes *[]models.Episode) {
	animeRepository.DB.Where("anime_id = ? ", id).Find(episodes)
}

func (animeRepository *AnimeRepository) GetAllAnimes(anime *[]models.Anime) {
	animeRepository.DB.Preload("Episodes").Find(anime)
}
