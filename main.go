/*
MIT License

Copyright (c) 2022 Georgios Kitsikoudis

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
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
