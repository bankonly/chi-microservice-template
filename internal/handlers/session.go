package handlers

import (
	"ecm-api-template/internal/models/dto"
	"ecm-api-template/internal/services"
	"net/http"

	"github.com/bankonly/go-pkg/v1/validator"
	"github.com/bankonly/go-pkg/v1/writers"
	"github.com/go-chi/chi/v5"
)

type Session interface {
	Route(chi.Router)                                  // All routers
	GenSession(w http.ResponseWriter, r *http.Request) // Microservice from another microservice
	GetSession(w http.ResponseWriter, r *http.Request) // Generate session for client access (Frontend)
}

type SessionOpts struct {
	services *services.Services
}

func NewSession() Session {
	return &SessionOpts{}
}

// Generate session for client access (Frontend)
func (opts *SessionOpts) GenSession(w http.ResponseWriter, r *http.Request) {
	writer := writers.Writer(w, r)

	var body dto.GenSessionRequestDTO
	if err := validator.Parser(r.Body, &body); err != nil {
		writer.BadRequest(err.Error())
		return
	}

	// Call session service to gen session
	res, err := opts.services.Session.GenSession(writer.RequestId(), &body)
	if err != nil {
		writer.ParseError(err)
		return
	}
	writer.JSON(res)
}

// Microservice from another microservice
func (opts *SessionOpts) GetSession(w http.ResponseWriter, r *http.Request) {
	writer := writers.Writer(w, r)

	// Call session service to get session
	opts.services.Session.GetSession(writer.RequestId())
	writer.Message("I am OK!")
}

// All routers
func (opts *SessionOpts) Route(r chi.Router) {
	r.Get("/", opts.GetSession)
	r.Get("/generate", opts.GenSession)
}
