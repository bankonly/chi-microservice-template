package middlewares

import (
	"ecm-api-template/internal/caches"
	"ecm-api-template/internal/configs"
	messageConf "ecm-api-template/internal/configs/message-conf"
	"ecm-api-template/pkg/storages"
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
)

func VerifySession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Development will be bypass
		if configs.Environment.MODE == configs.AppConf.DevLabel {
			next.ServeHTTP(w, r)
			return
		}

		writer := writers.New(w, r)
		vector := r.Header.Get("iv")
		token := r.Header.Get("authorization")
		enk := r.Header.Get("enk")

		if vector == "" || token == "" || enk == "" {
			writer.Forbidden(messageConf.ErrSessionDenied1)
			return
		}

		sessionMem := caches.NewSession(storages.GetRedis())
		if existedVector := sessionMem.GetVector(writer.RequestId(), vector); existedVector == vector {
			writer.Forbidden(messageConf.ErrSessionDenied2)
			return
		}

		// Save request vector
		if err := sessionMem.SaveVector(writer.RequestId(), vector); err != nil {
			writer.InternalServerError(messageConf.ErrSessionDenied3)
			return
		}
		next.ServeHTTP(w, r)
	})
}
