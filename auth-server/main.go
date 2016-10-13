package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Storage for user session tokens
var userStorage = make(map[string]string)

var existingUsers = []user{
	user{
		Username: "user1",
		Password: "pass1",
	},
	user{
		Username: "user2",
		Password: "pass2",
	},
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/login", login)
	// router.POST("/authenticate", authenticate)
	// router.POST("/logout", logout)

	router.Run(":13001")
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if userInDatabase(username, password) {
		token := generateToken()
		userStorage[username] = token
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func generateToken() string {
	// insecure way of token-generation
	return strconv.FormatInt(rand.Int63(), 16)
}

func userInDatabase(username, password string) bool {
	// Is user with given creds present in the database
	for _, u := range existingUsers {
		if username == u.Username && password == u.Password {
			return true
		}
	}
	return false
}
