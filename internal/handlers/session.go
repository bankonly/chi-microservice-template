package handlers

import (
	"ecm-api-template/internal/services"
	"net/http"

	"github.com/bankonly/go-pkg/v1/writers"
	"github.com/go-chi/chi/v5"
)

type Session interface {
	Route(chi.Router)                                  // All routers
	GenSession(w http.ResponseWriter, r *http.Request) // Microservice from another microservice
	GetSession(w http.ResponseWriter, r *http.Request) // Generate session for client access (Frontend)
}

type SessionOpts struct {
	services *services.Services // All services
}

func NewSession(services *services.Services) Session {
	return &SessionOpts{services: services}
}

// Generate session for client access (Frontend)
func (opts *SessionOpts) GenSession(w http.ResponseWriter, r *http.Request) {
	writer := writers.New(w, r)

	vector := r.Header.Get("iv")
	sessionId := r.Header.Get("enk-session")
	enk := r.Header.Get("enk")

	// Call session service to gen session
	data, err := opts.services.Session.GenSession(writer.RequestId(), vector, sessionId, enk)
	if err != nil {
		writer.ParseError(err)
		return
	}

	writer.JSON(data)
}

// Microservice from another microservice
func (opts *SessionOpts) GetSession(w http.ResponseWriter, r *http.Request) {
	writer := writers.New(w, r)

	// Call session service to get session
	opts.services.Session.GetSession(writer.RequestId())
	writer.Message("I am OK!")
}

// All routers
func (opts *SessionOpts) Route(r chi.Router) {
	r.Get("/", opts.GetSession)
	r.Post("/generate", opts.GenSession)
}
