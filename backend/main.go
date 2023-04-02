package main

import (
	"database/sql"
	_ "encoding/json"
	_ "net/http"
	"os"

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
	ID	uint
	Content	string
	Title	string
	Explain string
	UserId uint	
	CreatedAt string
	UpdatedAt string

}

func main() {
	r := gin.Default()
	// db := infrastructure.DbConn()
	dsn := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn), pgdriver.WithPassword("postgres")))
	db := bun.NewDB(sqldb, pgdialect.New())

	r.GET("/article", func(ctx *gin.Context) {
		var articles []Article
		err := db.NewSelect().Model((*Article)(nil)).Scan(ctx, &articles)
		if err != nil {
			util.ErrorNotice(err)
		}
		ctx.JSON(200, articles)
	})
	r.POST("/article", func (ctx *gin.Context) {
		//newArticle := 
	})

	godotenv.Load()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
