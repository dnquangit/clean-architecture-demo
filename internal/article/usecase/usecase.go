package usecase

import (
	"context"
	"go-module/internal/article"
	"go-module/internal/article/model"
)

type articleUC struct {
	repo article.Repo
}

func NewArticleUC(postRepo article.Repo) *articleUC {
	return &articleUC{
		repo: postRepo,
	}
}

func (us *articleUC) ListArticles(ctx context.Context, filter *model.Filter) (*[]model.Article, error) {
	return us.repo.ListArticles(ctx, filter)
}
