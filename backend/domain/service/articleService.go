package service

import (
	"context"

	"domain/model"
	"domain/repository"

)

type IArticleRepository interface {
	FindAllArticle(ctx context.Context) (model.ArticleSlice, error)
}

type articleRepository struct {
	repo repository.IArticleRepository
}

func (as *articleRepository) FindAllArticle(ctx context.Context) (model.ArticleSlice, error) {
	return as.repo.FindAllArticle(ctx)
}