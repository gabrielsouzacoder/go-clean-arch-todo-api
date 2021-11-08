package controllers

import (
	"github.com/gabrielsouzacoder/clean-new/api/presenter"
	"github.com/gabrielsouzacoder/clean-new/entity"
	"github.com/gabrielsouzacoder/clean-new/usecase/todo"
	"github.com/gin-gonic/gin"
)

func ListTodos(service *todo.Service) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		data, _ := service.ListTodos()

		c.JSON(200, data)
	}

	return fn
}

func DeleteTodo(service *todo.Service) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		id, _ := entity.StringToID(c.Param("id"))
		service.DeleteTodo(&id)

		c.Status(200)
	}

	return fn
}

func CreateTodo(service *todo.Service) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var todoDto presenter.Todo

		err := c.ShouldBindJSON(&todoDto)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "cannot bind JSON: " + err.Error(),
			})
			return
		}

		created, err := service.CreateTodo(todoDto.Description)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "cannot create todoDto: " + err.Error(),
			})
			return
		}

		c.JSON(200, created)
	}

	return fn
}

func FindById(service *todo.Service) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := entity.StringToID(c.Param("id"))

		todoById := service.FindById(&id)

		c.JSON(200, todoById)
	}

	return fn
}
