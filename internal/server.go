package internal

import (
	"ecm-api-template/internal/configs"
	"ecm-api-template/internal/routers"
	"ecm-api-template/pkg/middlewares"
	"log"
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServer() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(writers.Middleware)
	r.Use(middlewares.Recover(configs.Environment.MODE == configs.AppConf.ProdLabel))
	r.NotFound(middlewares.NotFound)
	r.MethodNotAllowed(middlewares.MethodNotAllow)

	r.Route("/v1/example-service", routers.Router)

	log.Println("Server is started on port", configs.Environment.PORT)
	http.ListenAndServe(":"+configs.Environment.PORT, r)
}
