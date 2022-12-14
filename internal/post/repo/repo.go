package repo

import (
	"context"
	"go-module/internal/post/model"
	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

func NewSQLRepo(db *gorm.DB) *postRepo {
	return &postRepo{db: db}
}

func (repo *postRepo) ListPosts(ctx context.Context, filter *model.Filter) (*[]model.Post, error) {
	return nil, nil
}
