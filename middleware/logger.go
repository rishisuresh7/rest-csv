package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type reqLogger struct {
	logger *logrus.Logger
}

func NewRequestLogger(l *logrus.Logger) Middleware {
	return &reqLogger{l}
}

func (l *reqLogger) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	startTime := time.Now()
	l.logger.WithFields(requestFields(r)).Infof("Request")

	next(w, r)

	fields := responseFields(r, w.(negroni.ResponseWriter))
	fields["Duration"] = int64(time.Since(startTime) / time.Millisecond)
	l.logger.WithFields(fields).Infof("Response")
}

func requestFields(r *http.Request) logrus.Fields {
	fields := logrus.Fields{
		"Client":     r.RemoteAddr,
		"Method":     r.Method,
		"URL":        r.URL.String(),
		"Referrer":   r.Referer(),
		"User-Agent": r.UserAgent(),
	}

	return fields
}

func responseFields(r *http.Request, w negroni.ResponseWriter) logrus.Fields {
	fields := logrus.Fields{
		"Method":     r.Method,
		"URL":        r.URL.String(),
		"StatusCode": w.Status(),
		"User-Agent": r.UserAgent(),
		"Size":       w.Size(),
	}

	return fields
}
