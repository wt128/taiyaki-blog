package auth0

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wt128/taiyaki-blog/util"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}
const ContextTokenKey = "token"
func NewMiddleware() (*jwtmiddleware.JWTMiddleware, error) {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			aud := os.Getenv("AUTH0_AUDIENCE")

			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAudience {
				return token, errors.New("Invalid audience.")
			}
			// verify iss claim
			iss := os.Getenv("AUTH0_API_AUDIENCE")
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := getPemCert(token)
			if err != nil {
				log.Fatalf("could not get cert: %+v", err)
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		UserProperty: ContextTokenKey,
		SigningMethod: jwt.SigningMethodRS256,
	}), nil
}

// var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("AUTH0_CLIENT_SECRET")), nil
// 	},
// 	SigningMethod: jwt.SigningMethodHS256,
// })
/*
func newValidationKeyGetter(domain, clientID string, jwks *JWKS) func(*jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return token, errors.New("invalid claims type")
		}
		azp, ok := claims["azp"].(string)

		if !ok {
			return nil, errors.New("authorized parties are required")
		}
		if azp != clientID {
			return nil, errors.New("invalid authorized parties")
		}

		iss := fmt.Sprintf("https://%s/", domain)
		ok = token.Claims.(jwt.MapClaims).VerifyIssuer(iss, true)
		if !ok {
			return nil, errors.New("invalid issuer")
		}

		cert, err := getPemCert(jwks, token)
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	}
} */

// jwksからjwtで使われているキーをpem形式で返す
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + os.Getenv("AUTH0_DOMAIN") + ".well-known/jwks.json")
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
		fmt.Print(jwtMiddleWare.Options.ValidationKeyGetter)
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
