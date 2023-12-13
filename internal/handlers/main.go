package handlers

import (
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
)

type Handlers struct {
	Session Session
}

func NewHandlers() *Handlers {
	return &Handlers{
		Session: NewSession(),
	}
}

func (opts *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writer := writers.Writer(w, r)

	writers.AssignLog(writer.RequestId(), map[string]interface{}{
		"index": "Server is running",
	})
	writer.Message("I am OK!")
}
