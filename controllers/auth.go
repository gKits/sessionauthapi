package controllers

import (
	"net/http"

	"github.com/gKits/sessionauthapi/database"
	"github.com/gKits/sessionauthapi/models"
	"github.com/gKits/sessionauthapi/token"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "failed to bind model"})
        return
    }

    user.EncryptPasswd()
    _, err := database.CreateUser(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "registration failed"})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

func Login(c *gin.Context) {
	var login models.Login

    session := sessions.Default(c)

	if err := c.BindJSON(&login); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "failed to bind model"})
		return
	}

	user, err := database.GetUserBy("Email", login.Email)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "user not found"})
		return
	}
    user.Password = login.Password

    err = database.ValidateUser(user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "wrong password"})
        return
    }

    token, err := token.CreatePasetoToken(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "token creation failed"})
        return
    }

    session.Set("username", user.Username)
    session.Set("token", token)
    err = session.Save(); if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func Logout(c *gin.Context) {
    session := sessions.Default(c)

    session.Delete("username")
    session.Delete("token")
    err := session.Save(); if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "session save failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func GetUser(c *gin.Context) {
    session := sessions.Default(c)

    username := session.Get("username")
    if username == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": nil, "message": "session corupted"})
        return
    }

    user, err := database.GetUserBy("Username", username.(string))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "user not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}









