package v1

import (
	"github.com/labstack/echo/v4"
)

// RegisterRoutes initializes api v1 routes
func (h *Handler) RegisterRoutes(e *echo.Group) {

	// Swagger routes
	e.GET("/v1/docs", h.getSwaggerIndex)
	e.GET("/v1/docs/swagger.yml", h.getSwaggerSchema)

	// persons retrieves all the persons.
	// swagger:route GET /persons persons persons
	//
	// Retreives all the persons.
	//
	// Produces:
	//  - application/json
	//
	// Security:
	//  - api_key:
	//
	// Responses:
	//  default: genericErrorResp
	//  200: personsSuccessResp
	//  400: badRequestErrorResp
	e.GET("/v1/persons", h.persons)

	// persons retrieves one person.
	// swagger:route GET /persons/:id persons person
	//
	// Retreives one person.
	//
	// Produces:
	//  - application/json
	//
	// Security:
	//  - api_key:
	//
	// Responses:
	//  default: genericErrorResp
	//  200: personSuccessResp
	//  400: badRequestErrorResp
	e.GET("/v1/persons/:id", h.person)

	// SavePerson register a new person in the database.
	// swagger:route POST /persons persons savePerson
	//
	// Registers a new person.
	//
	// Produces:
	//  - application/json
	//
	// Security:
	//  - api_key:
	//
	// Responses:
	//  default: genericErrorResp
	//  201: genericSuccessResp
	//  400: badRequestErrorResp
	e.POST("/v1/persons", h.savePerson)
}
