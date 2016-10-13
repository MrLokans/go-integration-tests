package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var auth = authService{Base: "http://localhost:4546"}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/login", loginFunc)
	router.GET("/logout", logoutFunc)
	router.POST("/protected-content", serveProtectedFunc)
	router.Run(":4545")
}

func loginFunc(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if response := auth.Login(username, password); response.Token != "" {
		c.SetCookie("username", username, 3600, "", "", false, true)
		c.SetCookie("token", response.Token, 3600, "", "", false, true)

		c.JSON(http.StatusOK, response)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func logoutFunc(c *gin.Context) {
	username, err1 := c.Cookie("username")
	token, err2 := c.Cookie("token")

	if err1 == nil && err2 == nil && auth.Logout(username, token) {
		c.SetCookie("username", "", -1, "", "", false, true)
		c.SetCookie("token", "", -1, "", "", false, true)

		c.JSON(http.StatusOK, nil)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func serveProtectedFunc(c *gin.Context) {
	username, err1 := c.Cookie("username")
	token, err2 := c.Cookie("token")

	if err1 == nil && err2 == nil && auth.Authenticate(username, token) {
		c.SetCookie("username", "", -1, "", "", false, true)
		c.SetCookie("token", "", -1, "", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "You are successfully authorized"})
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
