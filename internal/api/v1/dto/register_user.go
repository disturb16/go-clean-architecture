package dto

// RegisterUserDto parameters definitions.
//
// swagger:parameters savePerson
type RegisterUserDto struct {

	// in:body
	// required:true
	Body struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
}
