package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	_ "log"
	_ "net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

type Article struct {
	ID        uint   `json:"id" bun:id",pk,autoincrement"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	Explain   string `json:"explain"`
	UserId    uint   `json:"userId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func main() {
	r := gin.Default()
	// db := infrastructure.DbConn()
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"DELETE",
			"PUT",
		}}))
	dsn := "postgres://postgres:@db:5432/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn), pgdriver.WithPassword("postgres")))
	db := bun.NewDB(sqldb, pgdialect.New())

	r.GET("/article", func(ctx *gin.Context) {
		var article []Article
		if err := db.Ping(); err != nil {
			util.ErrorNotice(err)
		}
		err := db.NewSelect().Model((*Article)(nil)).Scan(ctx, &article)
		fmt.Println(article)
		if err != nil {
			util.ErrorNotice(err)
		}
		ctx.JSON(200, article)
		//log.Fatal(articles)
	})
	r.GET("/article/:id", func(ctx *gin.Context) {

	});
	r.POST("/article", func(ctx *gin.Context) {
		//newArticle :=
		title, _ := ctx.GetPostForm("title")
		content, _ := ctx.GetPostForm("content")
		userId, _ := ctx.GetPostForm("userId")
		intUserId, _ := strconv.Atoi(userId)
		newArticle := &Article{
			Title: title,
			UserId: uint(intUserId),
			Content: content,
		}
		if _, err := db.NewInsert().Model(newArticle).Exec(ctx); err != nil {
			util.ErrorNotice(err)
		}
		ctx.JSON(200, "success")
	})

	godotenv.Load()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
