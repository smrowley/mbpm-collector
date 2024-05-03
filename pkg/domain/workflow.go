package domain

import (
	"github.com/google/uuid"
)

type Workflow interface {
	GetId() uuid.UUID
}

type workflow struct {
	id          uuid.UUID
	description string
}

func (p *workflow) GetId() uuid.UUID {
	return p.id
}

func NewWorkflow(description string) Workflow {
	uuidVal, _ := uuid.NewRandom()

	workflow := workflow{
		id:          uuidVal,
		description: description,
	}

	return &workflow
}
