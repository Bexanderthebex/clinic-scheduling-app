package hospitals

type CreateHospitalReq struct {
	HospitalName string  `json:"name" validate:"required"`
	City         string  `json:"city" validate:"required"`
	Address      string  `json:"address" validate:"required"`
	Latitude     float64 `json:"latitude" validate:""`
	Longitude    float64 `json:"longitude" validate:""`
}

type FindHospitalReq struct {
	ID           string `json:"id" validate:"required_without_all=HospitalName"`
	HospitalName string `json:"name" validate:"required_without_all=ID"`
}
