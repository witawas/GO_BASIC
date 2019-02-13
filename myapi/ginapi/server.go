package main

import (
	"net/http"
	"strconv"

	"github.com/iamruchy/myapi/todo"

	"github.com/gin-gonic/gin"
)

//var todos = []todo.Todo{
//	todo.Todo{ID: "1", Title: "Test", Status: "active"},
//}

var todos = make(map[string]todo.Todo)

var todosArray = []todo.Todo{}

//var todos = []todo.Todo

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func jowHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Jow",
	})
}

func createJowHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Jow Post",
	})
}

func getTodosHandler(c *gin.Context) {
	id := c.Query("id")

	if id != "" {
		c.JSON(http.StatusOK, todos[id])
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func postTodosHandler(c *gin.Context) {
	var t todo.Todo
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&t); err != nil { //ShouldBindJSON
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	num := len(todos)
	t.ID = strconv.Itoa(num + 1)
	t.Status = "active"

	todos[t.ID] = t

	todosArray = append(todosArray, t)
	c.String(http.StatusOK, "Create OK")
}

func getTodosByIDHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, todos[id])
}

func updateTodosHandler(c *gin.Context) {
	id := c.Query("id")

	//c.JSON(http.StatusOK, todos[id])
	//return
	_, ok := todos[id]
	if ok {
		var t todo.Todo
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&t); err != nil { //ShouldBindJSON
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todos[id] = t
		c.String(http.StatusOK, "update complete")
		return
	} else {
		c.String(http.StatusOK, "No data found!!")
	}

}

func deleteTodosHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.String(http.StatusOK, "Please input id for delete")
	}
	//c.JSON(http.StatusOK, todos[id])
	//return
	_, ok := todos[id]
	if ok {
		delete(todos, id)
		c.String(http.StatusOK, "delete complete")
		return
	} else {
		c.String(http.StatusOK, "No data found!!")
	}

}

func setUp() *gin.Engine { // for test case
	r := gin.Default()
	v1 := r.Group("/v1")

	v1.GET("/hello", helloHandler)
	v1.GET("/jow", jowHandler)
	v1.POST("/jow", createJowHandler)
	v1.GET("/todos", getTodosHandler)

	v1.POST("/todos", postTodosHandler)

	v1.PUT("/todos", updateTodosHandler)

	v1.DELETE("/todos", deleteTodosHandler)

	v1.GET("/todos/:id", getTodosByIDHandler)

	return r
}

func main() {
	//fmt.Println("Start...")
	// r := gin.Default()
	// r.GET("/hello", helloHandler)
	// r.GET("/jow", jowHandler)
	// r.POST("/jow", createJowHandler)
	// r.GET("/todos", getTodosHandler)

	// r.POST("/todos", postTodosHandler)

	// r.PUT("/todos", updateTodosHandler)

	// r.DELETE("/todos", deleteTodosHandler)

	// r.GET("/todos/:id", getTodosByIDHandler)
	r := setUp()
	r.Run(":9082") // listen and serve on 0.0.0.0:8080
}
