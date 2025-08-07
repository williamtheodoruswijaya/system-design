package request

import "time"

type CreateTaskRequest struct {
	Title       string 		`json:"title"`
	Description *string 	`json:"description"`
	DueDate 	time.Time 	`json:"due_date"`
}