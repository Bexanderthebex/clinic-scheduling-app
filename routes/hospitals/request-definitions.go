package hospitals

type CreateHospitalReq struct {
	HospitalName string  `json:"name" validate:"required"`
	City         string  `json:"city" validate:"required"`
	Address      string  `json:"address" validate:"required"`
	Latitude     float64 `json:"latitude" validate:""`
	Longitude    float64 `json:"longitude" validate:""`
}

type SearchHospitalReq struct {
	HospitalName string `json:"name" validate:"required=ID"`
}
