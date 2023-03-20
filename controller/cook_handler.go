package controller

import (
	"fmt"
	"net/http"
	"strings"

	"cook/models"

	"github.com/gin-gonic/gin"
)

type CreateCookInput struct {
	Name       string `json:"name" binding:"required"`
	Ingredient string `json:"ingredient" binding:"required"`
}

type UpdateTaskInput struct {
	Name       string `json:"name" binding:"required"`
	Ingredient string `json:"ingredient" binding:"required"`
}

func FindTasks(c *gin.Context) {
	var tasks []models.Cook
	models.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateCookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	task := models.Cook{Name: input.Name, Ingredient: input.Ingredient}
	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func FindTask(c *gin.Context) { // Get model if exist
	var task models.Cook
	id := c.Request.URL.Query().Get("id")
	if err := models.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func UpdateTask(c *gin.Context) {
	// Get model if exist
	var task models.Cook
	id := c.Request.URL.Query().Get("id")
	if err := models.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func FindCookByName(c *gin.Context) { // Get model if exist
	var cook []models.Cook
	id := c.Request.URL.Query().Get("ingredient")
	res := strings.Replace(id, ",", "%%", -1)

	if err := models.DB.Where("ingredient LIKE ?", fmt.Sprintf("%%%s%%", res)).Find(&cook).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cook})

}
