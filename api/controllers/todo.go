package controllers

import (
	"fmt"
	"github.com/gabrielsouzacoder/clean-new/api/presenter"
	"github.com/gabrielsouzacoder/clean-new/usecase/todo"
	"github.com/gin-gonic/gin"
)

func ListTodos(service *todo.Service) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		data, _ := service.ListTodos()

		fmt.Println(data)

		c.JSON(200, data)
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
