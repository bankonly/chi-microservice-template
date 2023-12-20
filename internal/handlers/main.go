package handlers

import (
	"ecm-api-template/internal/services"
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
)

type Handlers struct {
	Session Session
}

func New(services *services.Services) *Handlers {
	return &Handlers{
		Session: NewSession(services),
	}
}

func (opts *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writer := writers.New(w, r)
	writers.AssignLog(writer.RequestId(), map[string]interface{}{
		"index": "Server is running",
	})
	writer.Message("I am OK!")
}
