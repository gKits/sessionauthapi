package middleware

import (
	"fmt"
	"net/http"

	"github.com/gKits/sessionauthapi/token"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthenticateSession(c *gin.Context) {
    abort := false
    session := sessions.Default(c)

    user := session.Get("username")
    t := session.Get("token")
    if user == nil || t == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": nil, "mesage": "session invalid"})
        abort = true
    } else {
        usernameString := user.(string)
        tokenString := t.(string)
        payload, err := token.ValidatePasetoToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(),"mesage": "token invalid"})
            abort = true
        } else if payload.Username != usernameString {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "session invalid","mesage": "session and token user not matching"})
            abort = true
        }
    }

    if abort {
        c.Abort()
    } else {
        fmt.Println("next")
        c.Next()
    }
}
