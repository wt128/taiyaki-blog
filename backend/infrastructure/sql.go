package infrastructure

import (
	"github.com/uptrace/bun"
)

type articleRepository struct {
	DB *bun.DB
}

/*
func (ar articleRepository) SelectAll(ctx context.Context) (model.ArticleSlice, error) {
	/* user := User{}
	err := db.NewSelect().Model(&user).Where("id = 1").Scan(context.Background())
	if err != nil {
		panic(err)
	}
	return model.Articles().All(ctx, ar.DB)
}  */
