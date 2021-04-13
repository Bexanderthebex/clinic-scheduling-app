package physician

import "fmt"

type PhysicianNotFoundError struct {
	PhysicianId string
}

func (p *PhysicianNotFoundError) Error() string {
	return fmt.Sprintf("Physician Id %s not found", p.PhysicianId)
}
