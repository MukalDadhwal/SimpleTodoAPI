package main

import (
	"fmt"
	"log"
	"os"

	"example/web-service-gin/config"
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}


func main() {
	if err := config.ConnectDB(); err != nil {
        log.Fatalf("❌ Failed to connect DB: %v", err)
    }

	defer config.DB.Close()

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	r := gin.Default()
    routes.RegisterRoutes(r)

	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
    r.Run(":" + port)
}
