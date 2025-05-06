package domain

import (
	"github.com/google/uuid"
)

func (id ProcessId) String() string {
	return uuid.UUID(id).String()
}

type Process interface {
	GetId() ProcessId
}

type process struct {
	id          ProcessId
	description string
}

func (p *process) GetId() ProcessId {
	return p.id
}

func NewProcess(description string) Process {
	uuidVal, _ := uuid.NewRandom()

	process := process{
		id:          ProcessId(uuidVal),
		description: description,
	}

	return &process
}
