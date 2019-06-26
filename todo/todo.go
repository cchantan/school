package todo
import (
	"school/schooldatabase"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
type Todohandler struct {}

func  (Todohandler) GetTodosHandler(c *gin.Context) {
	db, err := schooldatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT  id, title, status FROM todos")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	todos := []Todo{}
	t := Todo{}
	i := 1
	for rows.Next() {
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
	c.JSON(200, todos)
	return
	//
	//
}

func (Todohandler) GetTodosByIdHandler(c *gin.Context) {
	idinput := c.Param("id")
	db, err := schooldatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	stmt, _ := db.Prepare("SELECT  id, title, status FROM todos WHERE id=$1")
	//rows, _ := stmt.Queryrow()

	row := stmt.QueryRow(idinput)
	t := Todo{}
	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	fmt.Println("Test for One row => ", "\n ID     = ", t.ID, "\n Title  = ", t.Title, "\n Status = ", t.Status)
	fmt.Println(t)
	c.JSON(200, t)
	return
}

func (Todohandler) PostTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(t)

	db, err := schooldatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
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
	c.JSON(201, t)
	//	c.JSON(201,id)
}

func (Todohandler) DeleteTodosByIdHandler(c *gin.Context) {
	todos := []Todo{}
	idinput := c.Param("id")

	fmt.Println(idinput)

	db, err := schooldatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	query := `
		DELETE FROM todos WHERE id =$1
	`
	db.QueryRow(query, idinput)
	fmt.Println(todos)
	c.JSON(200, gin.H{
		"status": "success",
	})
	return
}

func (Todohandler) PutTodosByIdHandler(c *gin.Context) {
	idinput := c.Param("id")

	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(t)
	//
	db, err := schooldatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	query := `
	UPDATE todos SET title = $2, status = $3  WHERE id=$1
	`
	db.QueryRow(query, idinput, t.Title, t.Status)

	t.ID, err = strconv.Atoi(idinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, t)
	return
}