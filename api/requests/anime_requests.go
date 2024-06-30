package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateOrUpdateAnime struct {
	Title string `json:"title" validate:"required" example:"Echo"`
	Src   string `json:"src" validate:"required" example:"Echo is nice!"`
}

func (createOrUpdateAnime CreateOrUpdateAnime) Validate() error {
	return validation.ValidateStruct(&createOrUpdateAnime,
		validation.Field(&createOrUpdateAnime.Title, validation.Required),
		validation.Field(&createOrUpdateAnime.Src, validation.Required),
	)
}

type CreateAnimeRequest struct {
	CreateOrUpdateAnime
}

type UpdateAnimeRequest struct {
	CreateOrUpdateAnime
}
