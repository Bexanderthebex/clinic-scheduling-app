package physicians

type PhysicianRequestBody struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name"`
}

type PhysicianSpecializationRequestBody struct {
	PhysicianId     string   `json:"physician_id" validation:"required"`
	Specializations []string `json:"specializations" validate:"required"`
}
