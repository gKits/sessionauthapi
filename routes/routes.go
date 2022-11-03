package routes

import (
	"github.com/gKits/sessionauthapi/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
    g.POST("/login", controllers.Login)
    g.POST("/register", controllers.Register)
}

func PrivateRoutes(g *gin.RouterGroup) {
    g.GET("/logout", controllers.Logout)
    g.GET("/user", controllers.GetUser)
}
