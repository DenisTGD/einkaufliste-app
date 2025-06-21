package main

import "github.com/gin-gonic/gin"
import "github.com/DenisTGD/einkaufsliste-backend/internal/auth"




func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.POST("/register", auth.RegisterHandler)

	router.Run(":8080")
}
