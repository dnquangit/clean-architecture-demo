package repo

import (
	"context"
	"go-module/internal/article/model"
	"gorm.io/gorm"
)

type articleRepo struct {
	db *gorm.DB
}

func NewSQLRepo(db *gorm.DB) *articleRepo {
	return &articleRepo{db: db}
}

func (repo *articleRepo) ListArticles(ctx context.Context, filter *model.Filter) (*[]model.Article, error) {
	var posts []model.Article
	db := repo.db
	if err := db.Table(model.Article{}.TableName()).Where("deleted = ?", false).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
