package middlewares

import (
	"errors"
	"net/http"
	"github.com/ShubhamNatekar/Go-Mysql-API/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}