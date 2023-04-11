package auth0

import (
	"context"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwtmiddleware.JWTMiddlewareをContextに格納するためのキー
type JWTMiddlewareKey struct{}

// JWTをContextに保存するキー
type JWTKey struct{}

// jwtmiddleware.JWTMiddlewareをリクエストのContextに格納するためのMiddleware
func WithJWTMiddleware(m *jwtmiddleware.JWTMiddleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {
			// リクエストのContextにJWTMiddlewareを格納する
			context.WithValue(ctx.Request.Context(), JWTMiddlewareKey{}, m)
			//ctx.Request.Context().WithContext()
			// 新しいContextを入れて次の処理に渡す
			//next.ServeHTTP(w, r.WithContext(ctx))
			ctx.Next()
	}
}

// JWT検証を行うためのmiddleware
func UseJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ContextからJWTMiddlewareを取得
		jwtm := ctx.Request.Context().Value(JWTMiddlewareKey{}).(*jwtmiddleware.JWTMiddleware)
		// リクエスト中のJWTを検証
		if err := jwtm.CheckJWT(ctx.Writer, ctx.Request); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusUnauthorized)
			return
		}
		// JWT検証後に、Contextのjwtm.Options.UserPropertyからパース済みのトークンを取得する
		if val := ctx.Request.Context().Value(jwtm.Options.UserProperty); val != nil {
			token, ok := val.(*jwt.Token)
			if ok {
				// リクエストのContextにJWTを保存する
				addedJWTctx := context.WithValue(ctx.Request.Context(), JWTKey{}, token)
				// 新しいContextを入れて次の処理に渡す
				ctx.Request.WithContext(addedJWTctx)
				ctx.Next()
				return
			}
		}
		ctx.Next()
	}
}

// Contextに埋め込まれたJWTを取得する
func GetJWT(ctx context.Context) *jwt.Token {
	rawJWT, ok := ctx.Value(JWTKey{}).(*jwt.Token)
	if !ok {
		return nil
	}
	return rawJWT
}
