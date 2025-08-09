package response

import "time"

type TaskResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      int        `json:"status"`
	IsActive    int        `json:"is_active"`
	DueDate     time.Time  `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}