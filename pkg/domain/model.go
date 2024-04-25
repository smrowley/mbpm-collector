package domain

import (
	"github.com/google/uuid"
)

type Process interface {
	GetId() uuid.UUID
}

type process struct {
	id uuid.UUID
	description string
}

func (p *process) GetId() uuid.UUID {
	return p.id
}

func NewProcess(description string) Process {
	uuidVal, _ := uuid.NewRandom()

	process := process {
		id: uuidVal,
		description: description,
	}

	return &process
}