package handlers

import (
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
)

type HandlersImpl struct {
}

func (impl *HandlersImpl) Index(w http.ResponseWriter, r *http.Request) {
	wt := writers.Writer(w, r)

	writers.AssignLog(wt.RequestId(), map[string]interface{}{
		"index": "Server is running",
	})

	wt.Message("I am OK!")
}

func NewHandlers() *HandlersImpl {
	return &HandlersImpl{}
}
