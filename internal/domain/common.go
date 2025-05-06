package domain

import (
	"time"

	"github.com/google/uuid"
)

type IterationId uuid.UUID
type ParticipantId uuid.UUID
type ProcessId uuid.UUID
type Timestamp *time.Time
