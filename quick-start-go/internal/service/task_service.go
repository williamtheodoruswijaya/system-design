package service

import (
	"context"
	"database/sql"
	"quick-start-go/internal/model"
	"quick-start-go/internal/model/request"
	"quick-start-go/internal/model/response"
	"quick-start-go/internal/repository"
	"time"
)

// struktur-nya kurang lebih sama kek repository-nya, tapi ini lebih ke service layer-nya
type TaskService interface {
	Create(ctx context.Context, request request.CreateTaskRequest) (*response.TaskResponse, error)
	GetById(ctx context.Context, id int) (*response.TaskResponse, error)
	constructCreateTask(request request.CreateTaskRequest) model.Task // ini private function, ga perlu di expose ke luar
	constructTaskResponse(task model.Task) response.TaskResponse // ini private function juga
}

type TaskServiceImpl struct {
	// Karena cara membuat transaction harus menggunakan sql.DB, tapi kita ga punya parameter-nya, nah kita bisa define di sini
	DB *sql.DB

	// Sekalian kita panggil repository-nya disini karena services ini butuh repository-nya (yg dipanggil interface-nya)
	TaskRepository repository.TaskRepository
}

func NewTaskService(db *sql.DB, taskRepository repository.TaskRepository) TaskService { // parameter harus diisi sesuai dengan ...Impl
	return &TaskServiceImpl{
		DB:             db,
		TaskRepository: taskRepository,
	}
}

// Method
func (s *TaskServiceImpl) Create(ctx context.Context, request request.CreateTaskRequest) (*response.TaskResponse, error) {
	// step 1: begin transaction-nya
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	// step 2: buat function defer buat rollback-nya (seandainya kalau operasi dari step 4 ini gagal, maka defer ini akan dijalankan ketika semua function ini selesai sehingga rollback akan dijalankan agar insert-nya tidak dijalankan)
	defer func() error {
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}()

	// step 3: convert request ke model
	task := s.constructCreateTask(request)

	// step 4: call repository-nya buat create task-nya
	id, err := s.TaskRepository.Create(ctx, tx, task)
	if err != nil {
		return nil, err
	}

	// step 5: commit transaction-nya
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// step 6: ambil id-nya dari repository-nya
	response, err := s.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	// step 7: return response-nya
	return response, nil
}

func (s *TaskServiceImpl) GetById(ctx context.Context, id int) (*response.TaskResponse, error) {
	// langsung panggil dari repository-nya
	data, err := s.TaskRepository.GetById(ctx, s.DB, id)
	if err != nil {
		return nil, err
	}

	// convert model ke response
	response := s.constructTaskResponse(*data)

	// return response-nya
	return &response, nil
}

func (s *TaskServiceImpl) constructCreateTask(request request.CreateTaskRequest) model.Task { // buat function private untuk convert request ke model
	return model.Task{
		ID: 			0,
		Title:       	request.Title,
		Description: 	request.Description,
		Status:	  		0,
		IsActive: 		1,
		DueDate:		request.DueDate,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		nil,
	}
}

func (s *TaskServiceImpl) constructTaskResponse(task model.Task) response.TaskResponse {
	return response.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		IsActive:    task.IsActive,
		DueDate:     task.DueDate,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}