// routes.go
package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/todos", func(c *gin.Context) {
		fmt.Println("gettings todos...")
	})
	r.POST("/todos", func(c *gin.Context){
		fmt.Println("adding todos...")
	})
}
