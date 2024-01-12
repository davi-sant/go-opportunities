package router

import "github.com/gin-gonic/gin"

// Inicia esculta na porta :8080
func Routes() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run()
}
