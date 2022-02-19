package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"status": "fixed",
		})
	})
	r.Run("0.0.0.0:8080")
}
