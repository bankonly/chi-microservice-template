package api

import (
	"ecm-api-themplate/api/handlers"
	"log"
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
	"github.com/go-chi/chi/v5"
)

func NewServer() {
	r := chi.NewRouter()

	r.Use(writers.Middleware)

	indexHandlers := handlers.NewHandlers()
	r.Get("/", indexHandlers.Index)

	log.Println("Server is started on port", 9090)
	http.ListenAndServe(":9090", r)
}
