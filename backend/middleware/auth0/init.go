package auth0

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wt128/taiyaki-blog/util"
)

func Config() gin.HandlerFunc {
	godotenv.Load()
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")

	jwks, err := FetchJWKS(domain)
	if err != nil {
		util.ErrorNotice(err)
	}
	// domain, clientID, 公開鍵を元にJWTMiddlewareを作成する
	jwtMiddleware, err := NewMiddleware(domain, clientID, jwks)
	if err != nil {
		util.ErrorNotice(err)
	}
	fmt.Println(jwtMiddleware)
	return WithJWTMiddleware(jwtMiddleware)
}
