package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID string
	Body   string
	User   string
}
