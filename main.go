package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var myTodos = []todos{
	{ID: "1", Title: "Verify project budget", Completed: false, Category: 1},
	{ID: "2", Title: "Watch the movie!", Completed: false, Category: 2},
	{ID: "3", Title: "Feed the cat", Completed: false, Category: 3},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodos)
	router.GET("/todos/:id", todoById)
	router.Run("localhost:8080")
}

type todos struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Category  int    `json:"categoryId"`
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, myTodos)
}

func createTodos(c *gin.Context) {
	var newTodo todos
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	myTodos = append(myTodos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
func getTodoById(id string) (*todos, error) {
	for i, t := range myTodos {
		if strings.Compare(t.ID, id) > -1 { //if these strings are equal, then my result will be 0.
			return &myTodos[i], nil
		}
	}
	return nil, errors.New("There is no to-do like that id")
}
func todoById(c *gin.Context) {
	id := c.Param("id")
	todos, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "There is no to-do like that id"})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)

}
