package routes

import (
	"github.com/gabrielsouzacoder/clean-new/api/controllers"
	"github.com/gabrielsouzacoder/clean-new/usecase/todo"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service *todo.Service) *gin.Engine {
	main := router.Group("api/v1")
	{
		todos := main.Group("todos")
		{
			todos.GET("/", controllers.ListTodos(service))
			todos.GET("/:id", controllers.FindById(service))
			todos.POST("/", controllers.CreateTodo(service))
			todos.DELETE("/:id", controllers.DeleteTodo(service))
		}
	}
	return router
}
