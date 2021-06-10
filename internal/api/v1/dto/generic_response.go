package dto

import "github.com/disturb16/go-clean-architecture/internal/persons/entity"

type errorProperty struct {
	Property    string   `json:"property"`
	Constraints []string `json:"constraints"`
}

type errorBody struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// Generic success
// swagger:response genericSuccessResp
type GenericSuccessResp struct {
	// in: body
	Body struct {
		Data interface{} `json:"data"`
	}
}

// Generic Error
// swagger:response genericErrorResp
type GenericErrorResp struct {
	// Error message
	// in: body
	Body struct {
		Errors []struct {
			errorBody
			Properties []errorProperty `json:"properties"`
		} `json:"errors"`
	}
}

// Error in the request payload
//
// swagger:response badRequestErrorResp
type BadRequestResp struct {
	// Error message
	// in: body
	Body struct {
		Errors []struct {
			errorBody
		} `json:"errors"`
	}
}

// Representation of persons response
// swagger:response personsSuccessResp
type PersonsSuccessResp struct {
	// in: body
	Body struct {
		Data []entity.Person `json:"data"`
	}
}

// Representation of a person response
// swagger:response personSuccessResp
type PersonSuccessResp struct {
	// in: body
	Body struct {
		Data entity.Person `json:"data"`
	}
}
