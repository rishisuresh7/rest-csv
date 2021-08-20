package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"rest-csv/response"
	"rest-csv/utility"
)

type authenticator struct {
	logger *logrus.Logger
	routes []string
	token  string
}

type JWTAuthenticator struct {
	logger *logrus.Logger
	secret []byte
}

func NewAuthenticationMiddleware(l *logrus.Logger, t string, r []string) Middleware {
	return &authenticator{
		logger: l,
		token:  t,
		routes: r,
	}
}

func NewJWTAuthenticator(l *logrus.Logger, s string) *JWTAuthenticator {
	return &JWTAuthenticator{logger: l, secret: []byte(s)}
}

func (a *authenticator) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if utility.CheckList(a.routes, r.URL.Path) {
		next(w, r)
		return
	}

	if a.token != r.Header.Get("Api-Key") {
		response.Error{Error: "forbidden"}.Forbidden(w)
		return
	}

	next(w, r)
}

func (j *JWTAuthenticator) Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(tokenString) != 2 {
			j.logger.Errorf("invalid jwt token")
			response.Error{Error: "forbidden"}.Forbidden(w)
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString[1], claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return j.secret, nil
		})

		if err != nil {
			j.logger.WithError(err).Errorf("unexpected error happened")
			response.Error{Error: "forbidden"}.Forbidden(w)
			return
		}

		if !token.Valid {
			j.logger.Errorf("invalid token")
			response.Error{Error: "forbidden"}.Forbidden(w)
			return
		}

		next(w, r)
	}
}
