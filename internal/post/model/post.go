package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id           uuid.UUID  `gorm:"column:id;"`
	Title        string     `gorm:"column:title;"`
	ShortContent string     `gorm:"column:short_content;"`
	Content      string     `gorm:"column:content;"`
	Published    bool       `gorm:"type:boolean;column:published;"`
	PublishedAt  *time.Time `gorm:"column:published_at;"`
	CreateAt     time.Time  `gorm:"column:created_at;"`
	UpdateAt     time.Time  `gorm:"column:updated_at;"`
	Deleted      bool       `gorm:"column:deleted;"`
}
