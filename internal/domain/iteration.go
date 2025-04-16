package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type InstanceId uuid.UUID
type Timestamp *time.Time

func (id InstanceId) String() string {
	return uuid.UUID(id).String()
}

type Instance interface {
	GetId() InstanceId

	Start(participant uuid.UUID, startTime time.Time) error
	Trace(participant uuid.UUID, timestamp time.Time) error
	Complete(participant uuid.UUID, completionTime time.Time) error
}

type trace struct {
	participant uuid.UUID
	traceTime   Timestamp
}

type instance struct {
	id             InstanceId
	workflow       WorkflowId
	startTime      Timestamp
	completionTime Timestamp
	traces         []trace
}

func (i *instance) GetId() InstanceId {
	return i.id
}

func (i *instance) Start(participant uuid.UUID, startTime time.Time) error {
	if i.startTime != nil {
		return fmt.Errorf("Instance %v is already started", i.id)
	}
	if i.completionTime != nil {
		return fmt.Errorf("Instance %v is already completed", i.id)
	}
	i.startTime = Timestamp(&startTime)
	return nil
}

func (i *instance) Complete(participant uuid.UUID, completionTime time.Time) error {
	if i.startTime != nil {
		return fmt.Errorf("Instance %v is not started yet", i.id)
	}
	if i.completionTime != nil {
		return fmt.Errorf("Instance %v is already completed", i.id)
	}
	i.completionTime = Timestamp(&completionTime)
	return nil
}

func (i *instance) Trace(participant uuid.UUID, timestamp time.Time) error {
	if i.startTime != nil {
		return fmt.Errorf("Instance %v is not started yet", i.id)
	}
	i.traces = append(i.traces, trace{
		participant: participant,
		traceTime:   Timestamp(&timestamp),
	})
	return nil
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
