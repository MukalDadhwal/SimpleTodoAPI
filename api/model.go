package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed" default:"false"`
	CreatedAt time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Log     string `json:"log"`
}

func RespondWithError(c *gin.Context, status int, msg string, log string) {	
    c.JSON(status, ErrorResponse{
		Status:  status,
		Message: msg,
        Log: log,
	})
}
