package handler

import (
	"context"
	"net/http"
	"quick-start-go/internal/model/request"
	"quick-start-go/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

// ini juga ngikutin struktur dari repository dan service-nya, tapi ini lebih ke handler layer-nya (yang dipanggil dari controller-nya)

type TaskHandler interface {
	Create(c *gin.Context) // parameter ini samain kek yang ada di API
}

type TaskHandlerImpl struct {
	// Karena kita mau manggil service-nya, kita butuh service-nya di sini
	TaskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) TaskHandler { // Ingat karena handler akan hit services-nya, jadi kita butuh parameter service-nya di sini
	return &TaskHandlerImpl{
		TaskService: taskService,
	}
}

// Method
func (h *TaskHandlerImpl) Create(c *gin.Context) {
	// step 1: ambil request dari body-nya
	var request request.CreateTaskRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "bad request",
		})
		return
	}

	// step 2: buat context-nya (biar timeout otomatis kalau ga ada response dari service-nya)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // ini buat cancel context-nya setelah 5 detik

	// step 3: call service-nya buat create task-nya
	response, err := h.TaskService.Create(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": "internal server error",
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"code": http.StatusCreated,
			"message": "success",
			"data": response,
		})
	}
}