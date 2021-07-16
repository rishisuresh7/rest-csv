package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"rest-csv/response"
	"rest-csv/utility"
)

type authenticator struct {
	logger *logrus.Logger
	routes []string
	token  string
}

func NewAuthenticationMiddleware(l *logrus.Logger, t string, r []string) Middleware {
	return &authenticator{
		logger: l,
		token:  t,
		routes: r,
	}
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
