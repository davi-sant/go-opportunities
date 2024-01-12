package router

import "github.com/gin-gonic/gin"

// Inicia esculta na porta :8080
func Routes() {
	router := gin.Default()

	router.GET("api/teste/v0.0.1/opportunists", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"error":   false,
			"message": "Accept:code",
		})
	})
	router.Run()
}
