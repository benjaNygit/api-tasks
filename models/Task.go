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

func (t *Task) Create() bool {
	return false
}

func (t *Task) Read() Task {
	return Task{Id: t.Id}
}

func (t *Task) Update() Task {
	return Task{Id: t.Id}
}

func (t *Task) Delete() bool {
	return false
}
