package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type IterationId uuid.UUID
type Timestamp *time.Time

func (id IterationId) String() string {
	return uuid.UUID(id).String()
}

type Iteration interface {
	GetId() IterationId

	Start(participant uuid.UUID, startTime time.Time) error
	Trace(participant uuid.UUID, timestamp time.Time) error
	Complete(participant uuid.UUID, completionTime time.Time) error
}

type trace struct {
	participant uuid.UUID
	traceTime   Timestamp
}

type iteration struct {
	id             IterationId
	process       ProcessId
	startTime      Timestamp
	completionTime Timestamp
	traces         []trace
}

func (i *iteration) GetId() IterationId {
	return i.id
}

func (i *iteration) Start(participant uuid.UUID, startTime time.Time) error {
	if i.startTime != nil {
		return fmt.Errorf("Iteration %v is already started", i.id)
	}
	if i.completionTime != nil {
		return fmt.Errorf("Iteration %v is already completed", i.id)
	}
	i.startTime = Timestamp(&startTime)
	return nil
}

func (i *iteration) Complete(participant uuid.UUID, completionTime time.Time) error {
	if i.startTime != nil {
		return fmt.Errorf("Iteration %v is not started yet", i.id)
	}
	if i.completionTime != nil {
		return fmt.Errorf("Iteration %v is already completed", i.id)
	}
	i.completionTime = Timestamp(&completionTime)
	return nil
}

func (i *iteration) Trace(participant uuid.UUID, timestamp time.Time) error {
	if i.startTime != nil {
		return fmt.Errorf("Iteration %v is not started yet", i.id)
	}
	i.traces = append(i.traces, trace{
		participant: participant,
		traceTime:   Timestamp(&timestamp),
	})
	return nil
}

func NewIteration(processId uuid.UUID) Iteration {
	uuidVal, _ := uuid.NewRandom()

	iteration := iteration{
		id:       IterationId(uuidVal),
		process: ProcessId(processId),
		traces:   make([]trace, 10),
	}

	return &iteration
}
