package middlewares

import (
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			writer := writers.Writer(w, r)
			if r := recover(); r != nil {
				err := r.(string)
				writer.Status(http.StatusInternalServerError)
				writer.Message(err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
