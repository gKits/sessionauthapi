package middleware

import (
	"net/http"

    "github.com/gKits/sessionauthapi/token"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthenticateSession(c *gin.Context) {
    session := sessions.Default(c)
    user := session.Get("Username")
    if user == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": nil, "message": "session unauthorized"})
        c.Abort()
        return
    }

    t := session.Get("token")
    if t == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": nil, "message": "session unauthorized"})
        c.Abort()
        return
    }

    payload, err := token.ValidatePasetoToken(t.(string))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "token invalid"})
        c.Abort()
        return
    }

    if payload.Username != user.(string) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "token and session not matching"})
        c.Abort()
        return
    }

    c.Next()
}
