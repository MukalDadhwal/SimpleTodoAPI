package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"example/web-service-gin/db"
	"example/web-service-gin/routes"
)

// Mock data structure for a Todo
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var mockDB sqlmock.Sqlmock

// Mock router setup
func setupRouter() *gin.Engine {
	// Mock the database connection
	var err error
	db.DB, mockDB, err = sqlmock.New()
	if err != nil {
		panic("Failed to create mock database")
	}

	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}

func TestGetTodos(t *testing.T) {
	r := setupRouter()

	// Mock database response for GetTodos
	mockDB.ExpectQuery("SELECT id, title, completed, created_at FROM todos").WillReturnRows(
		sqlmock.NewRows([]string{"id", "title", "completed", "created_at"}).AddRow(1, "Test Todo", false, time.Now()),
	)

	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Debug log
	if w.Code != http.StatusOK {
		t.Logf("Response Body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your response structure
}

func TestCreateTodo(t *testing.T) {
	r := setupRouter()

	// Mock database response for PostTodos
	mockDB.ExpectQuery(`(?i)INSERT INTO todos \(title, completed, created_at\) VALUES \(\$1, \$2, \$3\) RETURNING id`).
		WithArgs("Test Todo", false, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	todo := Todo{
		Title:       "Test Todo",
		Description: "This is a test todo",
	}
	jsonValue, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Debug log
	if w.Code != http.StatusCreated {
		t.Logf("Response Body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	// Add more assertions based on your response structure
}

func TestUpdateTodo(t *testing.T) {
	r := setupRouter()

	// Mock database response for UpdateTodoPut
	mockDB.ExpectExec(`(?i)UPDATE todos SET title=\$1, completed=\$2 WHERE id=\$3`).
		WithArgs("Updated Todo", false, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	todo := Todo{
		Title:       "Updated Todo",
		Description: "This is an updated test todo",
	}
	jsonValue, _ := json.Marshal(todo)

	req, _ := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Debug log
	if w.Code != http.StatusNoContent {
		t.Logf("Response Body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusNoContent, w.Code)
	// Add more assertions based on your response structure
}

func TestDeleteTodo(t *testing.T) {
	r := setupRouter()

	// Mock database response for DeleteTodo
	mockDB.ExpectExec(`(?i)DELETE FROM todos WHERE id = \$1`).
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, _ := http.NewRequest("DELETE", "/todos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Debug log
	if w.Code != http.StatusNoContent {
		t.Logf("Response Body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusNoContent, w.Code)
	// Add more assertions based on your response structure
}
