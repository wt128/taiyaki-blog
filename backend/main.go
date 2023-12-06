package main

import (
	_ "encoding/json"
	_ "log"
	_ "net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/wt128/taiyaki-blog/infrastructure/db"
	"github.com/wt128/taiyaki-blog/middleware/auth0"
	corsMiddleware "github.com/wt128/taiyaki-blog/middleware/cors"
	"github.com/wt128/taiyaki-blog/util"
)

type AuthUser struct {
	ID           uint `bun:",pk,autoincrement"`
	Name         string
	Email        string
	Password     string
	CreatedAt    string `bun:created_at`
	UpdatedAt    string `bun:updated_at`
	Introduction string
}

type ArticleDTO struct {
	bun.BaseModel `bun:"table:articles,alias:ac"`
	ID            uint   `json:"id" bun:"id,pk,autoincrement"`
	Content       string `json:"content" bun:"content"`
	Tag           string `json:"tag" bun:"tag"`
	Title         string `json:"title" bun:"title"`
	CreatedAt     string `json:"createdAt" bun:"created_at"`
	Author        string `json:"author" bun:"author"`
}

type ArticleTag struct {
	bun.BaseModel `bun:"alias:at"`
	ID            uint   `json:"id" bun:"id,pk,autoincrement"`
	Name          string `json:"name"`
	Color         string `json:"color"`
}

func main() {
	r := gin.Default()
	// db := infrastructure.DbConn()
	r.Use(corsMiddleware.Config())
	//r.Use(auth0Middleware.Config())
	var db db.DB
	sqlInstance := db.DbConn()

	r.GET("/article", func(ctx *gin.Context) {
		var article []ArticleDTO
		err := sqlInstance.NewSelect().
			Model((*ArticleDTO)(nil)).
			ColumnExpr("ac.id ,title, content, u.name as author, ac.created_at, user_id, tag").
			Join("left join article_tags as at").
			JoinOn("ac.id = at.article_id").
			Join("left join auth_users as u").
			JoinOn("u.id = ac.user_id").
			Scan(ctx, &article)
		if err != nil {
			util.ErrorNotice(err)
			ctx.Abort()
		}
		ctx.JSON(200, article)
	})
	r.GET("/article/:id", func(ctx *gin.Context) {
		var article ArticleDTO
		
		id := ctx.Param("id")
		err := sqlInstance.NewSelect().
			Model((*ArticleDTO)(nil)).
			ColumnExpr("ac.id ,title, content, u.name as author, ac.created_at, user_id").
			Join("left join auth_users as u").
			JoinOn("u.id = ac.user_id").
			Where("ac.id = ?", id).
			Scan(ctx, &article)
		sqlInstance.NewSelect().
			Model((*ArticleTag)(nil)).
			Where("at.article_id = ?", id)


		if err != nil {
			util.ErrorNotice(err)
			ctx.Abort()
		}
		ctx.JSON(200, article)
	})

	r.POST("/article", auth0.AuthMiddleware(), HandlePost)
	godotenv.Load()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}

func HandlePost(ctx *gin.Context) {
	var db db.DB
	sqlInstance := db.DbConn()
	title, _ := ctx.GetPostForm("title")
	content, _ := ctx.GetPostForm("content")
	userId, _ := ctx.GetPostForm("userId")
	intUserId, _ := strconv.Atoi(userId)
	newArticle := map[string]interface{}{
		"title":   title,
		"user_id": uint(intUserId),
		"content": content,
	}
	if _, err := sqlInstance.NewInsert().Model(&newArticle).Table("articles").Exec(ctx); err != nil {
		util.ErrorNotice(err)
	}
	ctx.JSON(200, "success")
}
