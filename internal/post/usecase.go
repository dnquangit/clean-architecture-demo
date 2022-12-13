package post

import (
	"context"
	"go-module/internal/post/model"
)

type UseCase interface {
	ListPosts(ctx context.Context, filter *model.Filter) (*[]model.Post, error)
}
