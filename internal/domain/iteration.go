package domain

import (
	"fmt"

	"github.com/google/uuid"
)

func (id IterationId) String() string {
	return uuid.UUID(id).String()
}

type Iteration interface {
	GetId() IterationId

	Start(participant ParticipantId, startTime Timestamp) error
	Trace(participant ParticipantId, timestamp Timestamp) error
	Complete(participant ParticipantId, completionTime Timestamp) error
}

type trace struct {
	participant ParticipantId
	traceTime   Timestamp
}

type iteration struct {
	id             IterationId
	process        ProcessId
	startTime      Timestamp
	completionTime Timestamp
	traces         []trace
}

func (i *iteration) GetId() IterationId {
	return i.id
}

func (i *iteration) Start(participant ParticipantId, startTime Timestamp) error {
	if i.startTime != nil {
		return fmt.Errorf("Iteration %v is already started", i.id)
	}
	if i.completionTime != nil {
		return fmt.Errorf("Iteration %v is already completed", i.id)
	}
	i.startTime = startTime
	return nil
}

func (i *iteration) Complete(participant ParticipantId, completionTime Timestamp) error {
	if i.startTime != nil {
		return fmt.Errorf("Iteration %v is not started yet", i.id)
	}
	if i.completionTime != nil {
		return fmt.Errorf("Iteration %v is already completed", i.id)
	}
	i.completionTime = completionTime
	return nil
}

func (i *iteration) Trace(participant ParticipantId, timestamp Timestamp) error {
	if i.startTime != nil {
		return fmt.Errorf("Iteration %v is not started yet", i.id)
	}
	i.traces = append(i.traces, trace{
		participant: participant,
		traceTime:   timestamp,
	})
	return nil
}

func NewIteration(processId ProcessId) Iteration {
	uuidVal, _ := uuid.NewRandom()

	iteration := iteration{
		id:      IterationId(uuidVal),
		process: processId,
		traces:  make([]trace, 10),
	}

	return &iteration
}
