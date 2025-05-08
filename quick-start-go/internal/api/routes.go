package api

import (
	"database/sql"
	"quick-start-go/internal/handler"
	"quick-start-go/internal/repository"
	"quick-start-go/internal/service"

	"github.com/gin-gonic/gin"
)

// step 1: define sebuah Handler sebagai struct yang akan digunakan untuk mengumpulkan semua handler yang ada dalam rest api kita.
type Handlers struct {
	// simpen semua handler yang ada disini
	TaskHandler handler.TaskHandler
}

func SetupRoutes(db *sql.DB) *gin.Engine { // SetupRoutes akan dipanggil di file main.go yang mana dia sendiri akan memanggil initHandler dan initRoutes
	return initRoutes(initHandler(db))
}

func initHandler(db *sql.DB) Handlers {
	// tugas kita disini adalah menginisialisasi semua repository, service, dan handler yang ada dalam aplikasi kita, dan mengembalikannya dalam bentuk struct Handlers
	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(db, taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	return Handlers{
		TaskHandler: taskHandler,
	}
}

func initRoutes(handler Handlers) *gin.Engine { // Disini kita akan tuliskan semua jenis routes yang ada dalam aplikasi kita.
	// 1. Inisialisasi router dari Gin
	router := gin.Default()

	// 2. Lakukan GROUPING agar kita tidak mengulang-ulang penamaan sebuah route (contoh dibawah menunjukkan bahwa setiap routes yang dibuat harus diawali dengan "/api/v1")
	api := router.Group("/api/v1")
	tasks := api.Group("/tasks")

	// 3. Tulis end-point routes-routes-nya disini
	tasks.POST("", handler.TaskHandler.Create)

	return router
}