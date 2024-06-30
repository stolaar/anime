package post

import "anime/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
