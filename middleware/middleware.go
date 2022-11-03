package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthenticateSession(c *gin.Context) {
    session := sessions.Default(c)
    uid := session.Get("uid")
    if uid == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": nil, "message": "session unauthorized"})
        c.Abort()
        return
    }

    c.Next()
}
