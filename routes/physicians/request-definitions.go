package physicians

type PhysicianRequestBody struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name"`
}

type PhysicianSpecializationRequestBody struct {
	Specializations []string `json:"specializations" validate:"required"`
}

type Hospitals struct {
	Name    string  `json:"hospital_name" validate:"required"`
	City    string  `json:"city" validate:"required"`
	Address string  `json:"address" validate:"required"`
	lat     float32 `json:"lat" validate:"required"`
	long    float32 `json:"long" validate:"required"`
}

type CreateHospitalAffiliations struct {
	HospitalIDs []string `json:"hospital_ids" validate:"required"`
}
