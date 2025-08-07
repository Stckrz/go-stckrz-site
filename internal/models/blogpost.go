package models

import (
	"gorm.io/gorm"
)

type Blogpost struct {
	gorm.Model
	PostTitle string
	PostBody string
}
