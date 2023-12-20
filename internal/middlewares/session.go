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
		writer := writers.New(w, r)
		vector := r.Header.Get("iv")
		sessionInfo := r.Header.Get("enk-session")
		enk := r.Header.Get("enk")

		if vector == "" || sessionInfo == "" || enk == "" {
			writer.Forbidden(messageConf.ErrCredentialEmpty)
			return
		}

		// Development will be bypass
		if configs.Environment.MODE == configs.AppConf.DevLabel {
			next.ServeHTTP(w, r)
			return
		}

		sessionMem := caches.NewSession(storages.GetRedis())
		if existedVector := sessionMem.GetVector(writer.RequestId(), vector); existedVector == vector {
			writer.Forbidden(messageConf.ErrSessionDuplicated)
			return
		}

		// Save request vector
		if err := sessionMem.SaveVector(writer.RequestId(), vector); err != nil {
			writer.InternalServerError(messageConf.ErrSessionInternalFailed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
