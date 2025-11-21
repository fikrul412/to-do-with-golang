package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
    // Pagination
    page, _ := strconv.Atoi(c.Query("page"))
    limit, _ := strconv.Atoi(c.Query("limit"))
    if page < 1 { page = 1 }
    if limit < 1 { limit = 10 }

    // Filter / Search
    search := c.Query("search")
    categoryID := c.Query("category_id")

    var todos []models.Todo
    query := initializers.DB

    // Search by title or description
    if search != "" {
        query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
    }

    // Filter by category
    if categoryID != "" {
        query = query.Where("category_id = ?", categoryID)
    }

    // Pagination
    offset := (page - 1) * limit
    result := query.Offset(offset).Limit(limit).Find(&todos)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": todos})
}


func GetTodoById(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func CreateTodo(c *gin.Context) {
	var body models.Todo

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	todo := models.Todo{
		Title:       body.Title,
		Description: body.Description,
		Completed:   false,
		CategoryID:  body.CategoryID,
		Priority:    body.Priority,
		DueDate:     body.DueDate,
	}

	if err := initializers.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var body models.Todo

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var todo models.Todo
	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	initializers.DB.Model(&todo).Updates(models.Todo{
		Title:       body.Title,
		Description: body.Description,
		CategoryID:  body.CategoryID,
		Priority:    body.Priority,
		DueDate:     body.DueDate,
	})

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if err := initializers.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func ToggleComplete(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	initializers.DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
