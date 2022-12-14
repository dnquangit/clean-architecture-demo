package article

import (
	"context"
	"go-module/internal/article/model"
)

type Repo interface {
	ListArticles(ctx context.Context, filter *model.Filter) (*[]model.Article, error)
}
