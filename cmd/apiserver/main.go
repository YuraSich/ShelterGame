package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/YuraSich/ShelterGame/app"

	"github.com/gin-gonic/gin"
)

var (
	userSet app.UserSet
)

func main() {
	userSet.Init(1 << 16)
	r := gin.Default()
	r.GET("/", showUsers)
	r.GET("/user/:login", getUser)
	r.GET("/user", addUser)
	r.Run()
}

func getUser(c *gin.Context) {
	login := c.Param("login")
	usr := userSet.FindByLogin(login)
	if usr == nil {
		c.String(http.StatusNotFound, "Login "+login+" DOES NOT exist")
		return
	}
	c.String(http.StatusOK, "Login "+login+" exists id =  "+usr.ID)
}

func addUser(c *gin.Context) {
	login := c.Query("login")
	if login == "" {
		c.String(http.StatusBadRequest, "ID Or Login is invalid")
		return
	}
	usr := userSet.FindByLogin(login)
	if usr == nil {
		prevID, err := c.Cookie("userID")
		if err != nil {
			usr.ID = prevID
		} else {
			usr.ID = fmt.Sprint(rand.Uint64())
			c.SetCookie("userID", usr.ID, 3600, "/", "localhost", false, true)
		}
		userSet.Append(usr.ID, login)
		c.String(http.StatusOK, "User"+login+" Added id = "+usr.ID)
	} else {
		//TODO Если чувк уже заходил и у него есть куки
	}
}

func showUsers(c *gin.Context) {
	cUsers := userSet.GetUsers()
	for i := 0; i < len(cUsers); i++ {
		c.JSON(200, gin.H{
			cUsers[i].ID: cUsers[i].Login,
		})
	}

}
