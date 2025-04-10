package model

import "gorm.io/gorm"

type Comic struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey"`
	Filename        string `gorm:"uniqueIndex"`
	PreviewImageKey string
	Progress        int
	TotalPages      int
	LastOpenedTime  string
	CreatedTime     string
	UpdatedTime     string
}
