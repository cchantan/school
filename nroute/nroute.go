package nroute
import (
	"school/todo"
	"school/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Nroute() *gin.Engine {
	s := todo.Todohandler {}
	r := gin.Default()

	// Middleware use for every functions
	r.Use(middleware.Authmiddleware)
	//
	r.GET("/api/todos", s.GetTodosHandler)
	r.GET("/api/todos/:id", s.GetTodosByIdHandler)
	r.POST("/api/todos/", s.PostTodosHandler)
	r.DELETE("/api/todos/:id", s.DeleteTodosByIdHandler)
	r.PUT("/api/todos/:id", s.PutTodosByIdHandler)
	return r
}