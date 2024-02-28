package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CategoryCode int       `json:"category_code"`
	PriorityCode int       `json:"priority_code"`
	CreatedAt    time.Time `json:"created_at"`
	Done         bool      `json:"done"`
	UserId       uuid.UUID `json:"user_id"`
}
