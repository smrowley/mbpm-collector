package domain

import (
	"time"

	"github.com/google/uuid"
)

type Instance interface {
	GetId() uuid.UUID

	Start()
	Trace(participant uuid.UUID, timestamp time.Time)
	Complete()
}

type trace struct {
	participant uuid.UUID
	traceTime   time.Time
}

type instance struct {
	id             uuid.UUID
	workflow       uuid.UUID
	startTime      time.Time
	completionTime time.Time
	traces         []trace
}

func (i *instance) GetId() uuid.UUID {
	return i.id
}

func (i *instance) Start() {
	i.startTime = time.Now().UTC()
}

func (i *instance) Complete() {
	i.completionTime = time.Now().UTC()
}

func (i *instance) Trace(participant uuid.UUID) {
	i.traces = append(i.traces, trace{
		participant: participant,
		traceTime:   time.Now().UTC(),
	})
}

func NewInstance(workflowId uuid.UUID) Instance {
	uuidVal, _ := uuid.NewRandom()

	instance := instance{
		id:       uuidVal,
		workflow: workflowId,
		traces:   make([]trace, 10),
	}

	return &instance
}
