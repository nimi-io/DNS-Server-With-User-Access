package routes

import (
	// "dns-users/controllers"

	"dns-user/controllers"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(incommingRoutes *gin.RouterGroup){
    authGroup := incommingRoutes.Group("/auth")

	go  authGroup.POST("/register", controllers.Signup)
	 authGroup.POST("/login",/* controllers.Login*/)

}