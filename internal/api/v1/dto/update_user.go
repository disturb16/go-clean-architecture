package dto

// UpdateUserDto parameters definitions.
//
// swagger:parameters updatePerson
type UpdateUserDto struct {

	// in:body
	// required:true
	Body struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
}
