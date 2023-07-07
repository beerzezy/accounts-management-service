package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type jwtMiddleware struct {
	SigningKey   string
	SkipperPaths []string
}

func NewJWTMiddleware(signingKey string, skipperPaths []string) jwtMiddleware {
	return jwtMiddleware{
		SigningKey:   signingKey,
		SkipperPaths: skipperPaths,
	}
}

func (s jwtMiddleware) Handler() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			for _, skipperPath := range s.SkipperPaths {
				if path == skipperPath {
					return true
				}
			}
			return false
		},
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(s.SigningKey),
	})
}
