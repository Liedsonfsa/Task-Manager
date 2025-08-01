package models

import "time"

type Task struct {
	ID int 	`json:"id"`
	Title string `json:"title" validate:"required"`
	Description	string `json:"description"`
	Status string `json:"status" validate:"required,oneof=peding in_progress completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}