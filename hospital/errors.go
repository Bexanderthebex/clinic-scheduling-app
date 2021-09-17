package hospital

import "fmt"

type HospitalAlreadyExistsError struct {
	HospitalID string
}

func (h *HospitalAlreadyExistsError) Error() string {
	return fmt.Sprintf("Hospital with ID #{h.HospitalID} already exists")
}

type HospitalDoesNotExistError struct {
	HospitalID string
}

func (h HospitalDoesNotExistError) Error() string {
	return "Hospital does not exist"
}
