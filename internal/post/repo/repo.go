package repo

import (
	"context"
	"go-module/internal/post/model"
)

type postRepo struct {
}

func NewSQLRepo() *postRepo {
	return &postRepo{}
}

func (repo *postRepo) ListPosts(ctx context.Context, filter *model.Filter) (*[]model.Post, error) {
	return nil, nil
}
