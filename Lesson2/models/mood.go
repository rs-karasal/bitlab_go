package models

import "gorm.io/gorm"

type Mood struct {
	gorm.Model
	Value  uint8  `json:"value"` // 1 - 10
	Note   string `json:"note"`
	UserID uint   `json:"-"`
}
