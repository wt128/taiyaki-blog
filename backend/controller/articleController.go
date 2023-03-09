package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wt128/taiyaki-blog/usecase"
)

type articleController struct {
	usecase usecase.IArticleUsecase
}

func NewArticleController(au usecase.IArticleUsecase) *articleController {
	return &articleController{
		usecase: au,
	}
}

func (controller *articleController) FindAllArticle(c *gin.Context) {
	
	article, err := controller.usecase.FindAllArticle(ctx)
	if err := c.ShouldBindJSON(&article); err != nil {
		
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, article)
}
