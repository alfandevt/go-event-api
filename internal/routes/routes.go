package routes

import "github.com/gin-gonic/gin"

func RegisterServer(server *gin.Engine) {
	// User Routes
	server.POST("/signup", signUp)
}
