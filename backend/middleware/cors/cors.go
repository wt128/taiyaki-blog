package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Config() gin.HandlerFunc {
	return cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"DELETE",
			"PUT",
		},
		AllowHeaders: []string{
			"authorization",
			"Content-Type",
		},
	})
}
