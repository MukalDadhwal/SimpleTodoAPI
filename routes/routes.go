package routes

import (
	"example/web-service-gin/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", api.WelcomeEndpoint)
	r.GET("/todos/:id", api.GetTodoById)
	r.GET("/todos", api.GetTodos)
	r.POST("/todos", api.PostTodos)
	r.PUT("/todos/:id", api.UpdateTodoPut)
	r.PATCH("/todos/:id", api.UpdateTodosPatch)
	r.DELETE("/todos/:id", api.DeleteTodo)
}
