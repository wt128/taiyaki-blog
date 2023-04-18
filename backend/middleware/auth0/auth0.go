package auth0

import (
	"errors"
	"fmt"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

func NewMiddleware(domain, clientID string, jwks *JWKS) (*jwtmiddleware.JWTMiddleware, error) {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: newValidationKeyGetter(domain, clientID, jwks),
		SigningMethod:       jwt.SigningMethodRS256,
	}), nil
}

// var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("AUTH0_CLIENT_SECRET")), nil
// 	},
// 	SigningMethod: jwt.SigningMethodHS256,
// })

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
}

// jwksからjwtで使われているキーをpem形式で返す
func getPemCert(jwks *JWKS, token *jwt.Token) (string, error) {
	cert := ""
	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		return "", errors.New("unable to find appropriate key")
	}
	return cert, nil
}
