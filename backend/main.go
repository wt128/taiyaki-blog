package main

import (
	_ "encoding/json"
	_ "log"
	_ "net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

type Article struct {
	ID        uint   `json:"id" bun:id",pk,autoincrement"`
	Content   string `json:"content" "bun:content"`
	Title     string `json:"title" bun:"title"`
	Explain   string `json:"explain" bun:"explain"`
	UserId    uint   `json:"userId" bun:"user_id"`
	CreatedAt string `json:"createdAt" bun:"created_at"`
	UpdatedAt string `json:"updatedAt" bun:"updated_at"`
}

func main() {
	r := gin.Default()
	// db := infrastructure.DbConn()
	r.Use(corsMiddleware.Config())
	//r.Use(auth0Middleware.Config())
	var db db.DB
	sqlInstance := db.DbConn()
	r.GET("/article", func(ctx *gin.Context) {
		var article []Article
		err := sqlInstance.NewSelect().Model((*Article)(nil)).Scan(ctx, &article)
		if err != nil {
			util.ErrorNotice(err)
		}
		ctx.JSON(200, article)
	})
	r.GET("/article/:id", func(ctx *gin.Context) {
		var article Article
		id := ctx.Param("id")
		err := sqlInstance.NewSelect().Model((*Article)(nil)).Where("id = ?", id).Scan(ctx, &article)
		if err != nil {
			util.ErrorNotice(err)
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
	//token := auth0.GetJWT(ctx.Request.Context())
	// token.Claimsをjwt.MapClaimsへ変換
	//claims := token.Claims.(jwt.MapClaims)
	// claimsの中にペイロードの情報が入っている
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
