package main

import (
	"cook/controller"
	"cook/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/tasks", controller.FindTasks)
	r.POST("/api/tasks", controller.CreateTask)
	r.GET("/api/tasks/one", controller.FindTask)
	r.PUT("/api/tasks/update", controller.UpdateTask)
	r.GET("/api/cook", controller.FindCookByName)
	r.Run()
}
