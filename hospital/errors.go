package hospital

import "fmt"

type HospitalAlreadyExistsError struct {
	HospitalID string
}

func (h *HospitalAlreadyExistsError) Error() string {
	return fmt.Sprintf("Hospital with ID #{h.HospitalID} already exists")
}
