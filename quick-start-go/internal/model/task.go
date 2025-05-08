package model

import "time"

// isi dari file dalam model itu literally cuman struct doang
type Task struct {
	ID          int    		`json:"id"`
	Title       string 		`json:"title"`
	Description *string 	`json:"description"`
	Status 		int   		`json:"status"`
	IsActive 	int			`json:"is_active"`
	DueDate 	time.Time 	`json:"due_date"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	*time.Time 	`json:"updated_at"`
}