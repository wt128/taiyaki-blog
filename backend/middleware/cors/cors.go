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
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	})
}
