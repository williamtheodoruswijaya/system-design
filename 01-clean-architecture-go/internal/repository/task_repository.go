package repository

import (
	"context"
	"database/sql"
	"quick-start-go/internal/model"
)

type TaskRepository interface {
	// Blueprint method-method yang ada di task_repository.go
	Create(ctx context.Context, tx *sql.Tx, task model.Task) (*int, error)
	// GetAll()
	GetById(ctx context.Context, db *sql.DB, id int) (*model.Task, error)
	// Update()
	// Delete()
}

type TaskRepositoryImpl struct {
	// Nah, ini isinya struct yang kita implementasikan dari interface TaskRepository (untuk saat ini kita kosongin, ini ada karena struct ini wajib ada)
}

func NewTaskRepository() TaskRepository {
	// Ini buat inisialisasi TaskRepository interface dan struct-nya
	return &TaskRepositoryImpl{}
}

// Method-method sesuai dengan TaskRepository interface
func (r *TaskRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, task model.Task) (*int, error) {
	// step 1: define query-nya
	query := `INSERT INTO tasks(title, description, due_date) VALUES($1, $2, $3) RETURNING id;`

	// step 2: execute query-nya
	row := tx.QueryRowContext(ctx, query, task.Title, task.Description, task.DueDate)
	// akses id-nya dari row yang kita dapet dari queryRowContext
	
	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	// step 3: return id dan error
	return &id, nil
}

// Notes: Tx (Transaction) hanya akan digunakan dalam CREATE, UPDATE, dan DELETE saja
// karena kita butuh transaction-nya, sedangkan untuk READ kita tidak butuh transaction-nya, jadi kita bisa menggunakan DB (database) saja

func (r *TaskRepositoryImpl) GetById(ctx context.Context, db *sql.DB, id int) (*model.Task, error) {
	// step 1: define query-nya
	query := `SELECT id, title, description, status, is_active, due_date, created_at, updated_at FROM tasks WHERE id = $1;`

	// step 2: execute query-nya
	row := db.QueryRowContext(ctx, query, id)
	
	// step 3: scan row-nya ke struct Task
	var task model.Task
	err := row.Scan(&task.ID, 
					&task.Title, 
					&task.Description, 
					&task.Status, 
					&task.IsActive, 
					&task.DueDate, 
					&task.CreatedAt, 
					&task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}