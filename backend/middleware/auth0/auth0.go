package auth0

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/wt128/taiyaki-blog/util"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

const ContextTokenKey = "token"

func NewMiddleware() (*jwtmiddleware.JWTMiddleware, error) {

	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			aud := os.Getenv("AUTH0_API_AUDIENCE")

			if !verifyAud(token.Claims.(jwt.MapClaims), aud) {
				return token, errors.New("Invalid audience.")
			}
			// verify iss claim
			iss := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}
			// verify iss claim

			cert, err := getPemCert(token)
			if err != nil {
				log.Fatalf("could not get cert: %+v", err)
			}
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	}), nil
}

// 現状のライブラリ関数(VerifyAudience)だとaudが[]のときに対応していないため
// https://github.com/golang-jwt/jwt/blob/148d71010923ca691c950db9846191800f498f8d/map_claims.go#L69
// ここを参考に実装
func verifyAud(m map[string]interface{}, cmp string) bool {
	var cs []string
	var aud []string
	aud = append(aud, cmp)
	aud = append(aud, "https://" + os.Getenv("AUTH0_DOMAIN") + "/userinfo")
	switch v := m["aud"].(type) {
	case string:
		cs = append(cs, v)
	case []string:
		cs = v
	case []interface{}:
		for _, a := range v {
			vs, ok := a.(string)
			if !ok {
				return false
			}
			cs = append(cs, vs)
		}
	
	}
	fmt.Println(cs)
	fmt.Println(aud)
	return reflect.DeepEqual(cs, aud)
}
// jwksからjwtで使われているキーをpem形式で返す
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + os.Getenv("AUTH0_DOMAIN") + "/.well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return cert, err
	}

	x5c := jwks.Keys[0].X5c
	for k, v := range x5c {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, errors.New("unable to find appropriate key")
	}

	return cert, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client secret key
		jwtMiddleWare, err1 := NewMiddleware()

		if err1 != nil {
			util.ErrorNotice(err1)
		}
		err2 := jwtMiddleWare.CheckJWT(c.Writer, c.Request)
		if err2 != nil {
			// Token not found
			fmt.Println(err2)
			util.ErrorNotice(err2)
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
	}
}
