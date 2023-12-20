package middlewares

import (
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	writer := writers.New(w, r)
	writer.NotFound("unimplemented")
}
