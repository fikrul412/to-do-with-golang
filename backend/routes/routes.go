package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/todos", controllers.GetTodos)
	api.POST("/todos", controllers.CreateTodo)
	api.GET("/todos/:id", controllers.GetTodoById)
	api.PUT("/todos/:id", controllers.UpdateTodo)
	api.DELETE("/todos/:id", controllers.DeleteTodo)
	api.PATCH("/todos/:id/complete", controllers.ToggleComplete)

	api.GET("/categories", controllers.GetCategories)
	api.POST("/categories", controllers.CreateCategory)
}
