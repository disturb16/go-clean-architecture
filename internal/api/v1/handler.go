package v1

import (
	"errors"

	"github.com/disturb16/go-sqlite-service/internal/persons"
	"github.com/disturb16/go-sqlite-service/settings"
	validator2 "github.com/sanservices/apicore/validator"
)

var errParametersNotValid = errors.New("one or more parameters are incorrectly formatted")

// Handler handles v1 routes
type Handler struct {
	cfg       *settings.Settings
	service   persons.Service
	validator *validator2.Validator
}

// NewHandler initialize main *Handler
func NewHandler(cfg *settings.Settings, svc persons.Service, validator *validator2.Validator) *Handler {

	return &Handler{
		cfg:       cfg,
		service:   svc,
		validator: validator,
	}
}
