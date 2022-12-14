package viewmodel

import (
	"github.com/google/uuid"
	"time"
)

type ArticleResponse struct {
	Id           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	ShortContent string    `json:"content"`
	UpdateAt     time.Time `json:"updated_at"`
}
