package main

import (
	"context"
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wt128/taiyaki-blog/domain/model"
	"github.com/wt128/taiyaki-blog/infrastructure"
	"github.com/wt128/taiyaki-blog/util"
)

type hoge2 struct {
	Id       int64  `bun:"id"`
	Name     string `bun:"name"`
	Age      int    `bun:"age"`
	Password string `bun:"password"`
}

func main() {
	r := gin.Default()
	// db := infrastructure.DbConn()
	r.GET("/ping", func(c *gin.Context) {
		db := infrastructure.DbConn()
		article := []model.Article{}
		err := db.NewSelect().Model(&article).Scan(context.Background())
		_, err1 := db.NewCreateTable().Model((*hoge2)(nil)).Exec(context.Background())
		if err1 != nil {
			util.ErrorNotice(err1)
		}
		if err != nil {
			util.ErrorNotice(err1)
		}
		//articleDTO, err := json.Marshal(article)
		if err != nil {
			c.AbortWithStatus(http.StatusUnprocessableEntity)
		}
		c.JSON(200, db.String())
	})
	r.Run(":8080")

	// Open a PostgreSQL database.
	/*	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
		pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

		// Create a Bun db on top of it.
		db := bun.NewDB(pgdb, pgdialect.New())

		// Print all queries to stdout.
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

		var rnd float64

		// Select a random number.
		if err := db.NewSelect().ColumnExpr("random()").Scan(ctx, &rnd); err != nil {
			panic(err)
		}

		fmt.Println(rnd) */
}
