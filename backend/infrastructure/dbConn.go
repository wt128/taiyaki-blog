package infrastructure

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/wt128/taiyaki-blog/util"
	_ "github.com/wt128/taiyaki-blog/util"
)

func DbConn() (*bun.DB) {
	// postgresql://{ホスト名}:{ポート番号}/{DB名}?user={ユーザ名}&password={パスワード}
  dsn := "postgresql://localhost:5432/postgres?user=postgres&password=postgres&sslmode=disable"
	db := bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgresql://localhost:5432/postgres?sslmode=disable"))),
		pgdialect.New(),
	)

	var v string
	if err := db.NewSelect().ColumnExpr("version()").Scan(context.Background(), &v); err != nil {
		util.ErrorNotice(err)
	}
	fmt.Println(v)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	return bun.NewDB(sqldb, pgdialect.New())
}
