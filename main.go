//package main
//import (
//	"fmt"
//	"net/http"  
//	"github.com/gin-gonic/gin"
//   )
//func getTodos(c *gin.Context) {/
//	c.JSON(200, "Okay")/
//}
//func main(){
//	r := gin.Default()
//	r.GET("/api/todos", getTodos)
//	r.Run(";1234")
//}
package main

import (
	"database/sql"
	"strconv"
	"os"
	"log"
 	"fmt"
 	"net/http"  
 	"github.com/gin-gonic/gin"
 _ 	"github.com/lib/pq"
)
type Todo struct {
	ID		int `json:"id"`
	Title	string `json:"title"`
	Status	string `json:"status"`
}

func getTodos(c *gin.Context) {
	c.JSON(200, "Okay")
}
func getTodosHandler(c *gin.Context){
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	stmt, _ := db.Prepare("SELECT  id, title, status FROM todos")
	rows, _ := stmt.Query()
	defer db.Close()

	todos := []Todo{}
	t := Todo{}
	i := 1
	for rows.Next(){
	//	t := Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}
		fmt.Println("Test for One row => ", i, "\nID     = ", t.ID, "\nTitle  = ", t.Title, "\nStatus = ", t.Status)
		todos = append(todos, t)
		i++
		
	}
	fmt.Println(todos)
	c.JSON(200,todos)
	return
	//
	//
}

func getTodosByIdHandler(c *gin.Context){
	idinput := c.Param("id")
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	stmt, _ := db.Prepare("SELECT  id, title, status FROM todos WHERE id=$1")
	//rows, _ := stmt.Queryrow()
	defer db.Close()

	row := stmt.QueryRow(idinput)
	t := Todo{}	
	err := row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	fmt.Println("Test for One row => ", "\n ID     = ", t.ID, "\n Title  = ", t.Title, "\n Status = ", t.Status)
	fmt.Println(t)
	c.JSON(200,t)
	return
}

func postTodosHandler(c *gin.Context){
	t := Todo{}	
	if err:=c.ShouldBindJSON(&t); err!=nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}
	fmt.Println(t)
	//
	url := os.Getenv("DATABASE_URL")
	fmt.Println("Url =", url)
	db, err := sql.Open("postgres", url)
	if err!= nil {
		log.Fatal("fatal", err.Error())
	}

	defer db.Close()

	title := t.Title
	status := t.Status
	query := `
		INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id
	`
	var id int 
	row := db.QueryRow(query, title, status)
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't scan id", id)
	}
	fmt.Println("Insert success id", id)
	t.ID = id
	c.JSON(201,t)
//	c.JSON(201,id)
}

func deleteTodosByIdHandler(c *gin.Context){
	todos := []Todo{}
	idinput := c.Param("id")
	
	fmt.Println(idinput)

	url := os.Getenv("DATABASE_URL")
	fmt.Println("Url =", url)
	db, err := sql.Open("postgres", url)
	if err!= nil {
		log.Fatal("fatal", err.Error())
	}

	defer db.Close()

	query := `
		DELETE FROM todos WHERE id =$1
	`
	db.QueryRow(query, idinput)
	fmt.Println(todos)
	c.JSON(200,gin.H{
		"status": "success",
	})
	return
}

func putTodosByIdHandler(c *gin.Context){
	idinput := c.Param("id")

	t := Todo{}	
	if err:=c.ShouldBindJSON(&t); err!=nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}
	fmt.Println(t)
	//
	url := os.Getenv("DATABASE_URL")
	fmt.Println("Url =", url)
	db, err := sql.Open("postgres", url)
	if err!= nil {
		log.Fatal("fatal", err.Error())
	}
	defer db.Close()

	query := `
	UPDATE todos SET title = $2, status = $3  WHERE id=$1
	`
	db.QueryRow(query,idinput, t.Title, t.Status)

	t.ID, err = strconv.Atoi(idinput)
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}
	c.JSON(200,t)
	return
}

func main(){
	r := gin.Default()
	r.GET("/api/todos", getTodosHandler)
	r.GET("/api/todos/:id", getTodosByIdHandler)
	r.POST("/api/todos/", postTodosHandler)
	r.DELETE("/api/todos/:id", deleteTodosByIdHandler)
	r.PUT("/api/todos/:id", putTodosByIdHandler)

	r.Run(":1234")
}