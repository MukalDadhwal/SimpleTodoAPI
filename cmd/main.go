package main

import (
	"fmt"
	"log"
	"os"

	"example/web-service-gin/api"
	"example/web-service-gin/db"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(r *gin.Engine) {
	r.GET("/", api.WelcomeEndpoint)
	r.GET("/todos", api.GetTodos)
	r.POST("/todos", api.PostTodos)
}

func main() {
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	defer db.DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	RegisterRoutes(r)

	fmt.Printf("Server running at http://localhost:%s\n", port)
	r.Run(":" + port)
}
