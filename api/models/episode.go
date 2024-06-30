package models

import "gorm.io/gorm"

type Episode struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(200);"`
	Description string `json:"description" gorm:"type:varchar(200);"`
	Link        string `json:"link" gorm:"type:varchar(200);"`
	Src         string `json:"src" gorm:"type:varchar(200);"`
	Image       string `json:"image" gorm:"type:varchar(200);"`
	AnimeID     uint
}
