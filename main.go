package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gKits/sessionauthapi/middleware"
	"github.com/gKits/sessionauthapi/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load(); if err != nil {
        log.Fatal("error loading .env file")
    }
}

func main() {
    router := gin.Default()

    router.Use(sessions.Sessions("sessions", cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))))

    public := router.Group("/api")
    routes.PublicRoutes(public)

    private := router.Group("/api")
    private.Use(middleware.AuthenticateSession)
    routes.PrivateRoutes(private)

    router.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
