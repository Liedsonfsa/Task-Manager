package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(taskHandler *TaskHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	tasks := router.Group("/tasks")
	{
		tasks.POST("", taskHandler.CreateTask)
        tasks.GET("", taskHandler.GetAllTasks)
        tasks.GET("/:id", taskHandler.GetTask)
        tasks.PUT("/:id", taskHandler.UpdateTask)
        tasks.DELETE("/:id", taskHandler.DeleteTask)
	}

	return router
}