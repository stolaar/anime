package post

import "anime/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
