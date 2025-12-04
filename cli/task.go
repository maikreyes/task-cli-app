package cli

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createAt"`
	UpdatedAt   time.Time `json:"UpdateAt"`
}
