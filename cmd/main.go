package main

import (
	"fmt"
	"log"
	"os"

	"example/web-service-gin/db"
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

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
	routes.RegisterRoutes(r)

	fmt.Printf("Server running at http://localhost:%s\n", port)
	r.Run(":" + port)
}
