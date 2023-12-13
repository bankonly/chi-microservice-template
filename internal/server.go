package internal

import (
	"ecm-api-template/internal/routers"
	"log"
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServer() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(writers.Middleware)

	r.Route("/public", routers.PublicRouter)

	log.Println("Server is started on port", 9090)
	http.ListenAndServe(":9090", r)
}
