package usecase

import (
	"context"
	"go-module/internal/post"
	"go-module/internal/post/model"
)

type postUC struct {
	repo post.Repo
}

func NewPostUC(postRepo post.Repo) *postUC {
	return &postUC{
		repo: postRepo,
	}
}

func (us *postUC) ListPosts(ctx context.Context, filter *model.Filter) (*[]model.Post, error) {
	return us.repo.ListPosts(ctx, filter)
}
