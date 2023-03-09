package repository

import (
	"content"
	"domain/model"
)

type IArticleRepository interface {
	SelectAll(ctx content.Context) (model.ArticleSlice, err);
}