package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id           uuid.UUID
	Title        string
	Description  string
	CategoryCode int
	PriorityCode int
	CreatedAt    time.Time
	Done         bool
	UserId       uuid.UUID
}
