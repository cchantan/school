package main

import (
	"school/todo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func main() {
	s := todo.Todohandler {}
	r := gin.Default()
	r.GET("/api/todos", s.GetTodosHandler)
	r.GET("/api/todos/:id", s.GetTodosByIdHandler)
	r.POST("/api/todos/", s.PostTodosHandler)
	r.DELETE("/api/todos/:id", s.DeleteTodosByIdHandler)
	r.PUT("/api/todos/:id", s.PutTodosByIdHandler)

	r.Run(":1234")
}
