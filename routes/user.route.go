package routes

import (
	// "dns-users/controllers"

	con "dns-user/controllers"
	mid "dns-user/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.RouterGroup) {
	authGroup := incommingRoutes.Group("")

	authGroup.GET("/user", mid.AuthMiddleware(), con.GetUser)
	//  authGroup.POST("/login", controllers.SignIn)

}
