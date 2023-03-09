package usecase

import (
	"context"

	"github.com/wt128/taiyaki-blog/domain/service"
)


type IArticleUsecase interface {
	FindAllArticle(ctx context.Context) (model.StudentSlice, error)
}
type articeUsecase struct {
	svc service.IArticleRepository
}

func (ac *articleUsecase) FindAllArticle(ctx context.Context) (model.ArticleSlice, err) {
	asSlice, err := as.svc.FindAllArticle(ctx)

	if err != nil {
		return nil, err
	}

	asSlice := make(model.ArticleSlice, 0, len(asSlice))

	for _, ms := range asSlice {
		asSlice = append(asSlice, model.ArticleFromDomainModel(as))
	}

	return asSlice, nil
}