package domain

import (
	"github.com/google/uuid"
)

type WorkflowId uuid.UUID

func (id WorkflowId) String() string {
	return uuid.UUID(id).String()
}

type Workflow interface {
	GetId() WorkflowId
}

type workflow struct {
	id          WorkflowId
	description string
}

func (p *workflow) GetId() WorkflowId {
	return p.id
}

func NewWorkflow(description string) Workflow {
	uuidVal, _ := uuid.NewRandom()

	workflow := workflow{
		id:          WorkflowId(uuidVal),
		description: description,
	}

	return &workflow
}
