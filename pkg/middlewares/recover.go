package middlewares

import (
	messageConf "ecm-api-template/internal/configs/message-conf"
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
)

func Recover(isProduction bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				writer := writers.New(w, r)
				if r := recover(); r != nil {
					err := r.(string)
					writer.Status(http.StatusInternalServerError)
					writers.SetError(writer.RequestId(), err)
					if !isProduction {
						writer.Message(err)
					} else {
						writer.Message(messageConf.ErrInternalServerError)
					}
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
