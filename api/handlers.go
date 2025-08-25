package api

import (
	"example/web-service-gin/db"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func WelcomeEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Todo API!"})
}

func GetTodos(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, completed, created_at FROM todos")

	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Cannot fetch the todos from the table", err.Error())
		return
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt); err != nil {
			RespondWithError(c, http.StatusInternalServerError, "Error fetch todos", err.Error())
			return
		}
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, todos)

	if err := rows.Err(); err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Error fetching rows", err.Error())
	}
}

func PostTodos(c *gin.Context) {
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO todos (title, completed, created_at) VALUES ($1, $2, $3) RETURNING id"
	err := db.DB.QueryRow(query, todo.Title, todo.Completed, time.Now()).Scan(&todo.ID)

	if err != nil{
		RespondWithError(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func GetTodoById(c *gin.Context){
	id := c.Param("id")

	var todo Todo

	row := db.DB.QueryRow("SELECT * FROM todos WHERE id = $1", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt); err != nil{	
		RespondWithError(c, http.StatusNotFound, "The todo does not exist", err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodoPut(c *gin.Context){
	id := c.Param("id")
	var todo Todo

	if err := c.BindJSON(&todo); err != nil{
		RespondWithError(c, http.StatusInternalServerError, "Internal Server error", err.Error())
		return
	}
	
	_, err := db.DB.Exec("UPDATE todos SET title=$1, completed=$2 WHERE id=$3", &todo.Title, &todo.Completed, id)

	if err != nil{
		RespondWithError(c, http.StatusInternalServerError, "Internal server error", err.Error())
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Todo updated"})
}

func UpdateTodosPatch(c *gin.Context){
	id := c.Param("id")

	var payload map[string]any

	if err := c.BindJSON(&payload); err != nil{
		RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return;
	}

	query := "UPDATE todos SET "
	params := []any{}
	i := 1

	if title, ok := payload["title"]; ok {
		query += "title = $" + strconv.Itoa(i) + ", "
		params = append(params, title)
		i++
	}

	if completed, ok := payload["completed"]; ok {
		query += "completed = $" + strconv.Itoa(i) + ", "
		params = append(params, completed)
		i++
	}

	query = query[:len(query) - 2]
	query += " WHERE id = $" + strconv.Itoa(i)
	params = append(params, id)

	fmt.Println(params)
	fmt.Println(len(params))
	fmt.Print(query)
	_, err := db.DB.Exec(query, params...)

	if err != nil{
		RespondWithError(c, http.StatusInternalServerError, "Failed to updated todo", err.Error())
		return
	}


	c.JSON(http.StatusNoContent, gin.H{"message": "Todo updated successfully"})

}

func DeleteTodo(c *gin.Context){
	id := c.Param("id")

	query := "DELETE FROM todos WHERE id = $1"
	_, err := db.DB.Exec(query, id)

	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Failed to delete todo", err.Error())
		return;
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "todo deleted successfully"})
}
