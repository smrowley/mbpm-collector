package domain

import (
	"time"

	"github.com/google/uuid"
)

type InstanceId uuid.UUID

type Instance interface {
	GetId() InstanceId

	Start(participant uuid.UUID, startTime time.Time)
	Trace(participant uuid.UUID, timestamp time.Time)
	Complete(participant uuid.UUID, completionTime time.Time)
}

type trace struct {
	participant uuid.UUID
	traceTime   time.Time
}

type instance struct {
	id             InstanceId
	workflow       WorkflowId
	startTime      time.Time
	completionTime time.Time
	traces         []trace
}

func (i *instance) GetId() InstanceId {
	return i.id
}

func (i *instance) Start(participant uuid.UUID, startTime time.Time) {
	i.startTime = startTime
}

func (i *instance) Complete(participant uuid.UUID, completionTime time.Time) {
	i.completionTime = completionTime
}

func (i *instance) Trace(participant uuid.UUID, timestamp time.Time) {
	i.traces = append(i.traces, trace{
		participant: participant,
		traceTime:   timestamp,
	})
}

func NewInstance(workflowId uuid.UUID) Instance {
	uuidVal, _ := uuid.NewRandom()

	instance := instance{
		id:       InstanceId(uuidVal),
		workflow: WorkflowId(workflowId),
		traces:   make([]trace, 10),
	}

	return &instance
}
