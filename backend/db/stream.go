package db

import "gorm.io/gorm"

type Streaming struct {
	gorm.Model
	UUID string ``
}
