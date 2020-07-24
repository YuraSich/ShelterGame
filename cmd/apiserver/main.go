package main

import (
	"encoding/json"
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

func userHandler(c *gin.Context) {

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
		c.String(http.StatusBadRequest, "Login is invalid")
		return
	}

	usr := userSet.FindByLogin(login)

	if usr != nil {
		c.String(http.StatusBadRequest, "User exists")
		return
	}

	newID := fmt.Sprint(rand.Uint64())
	for userSet.FindByID(newID) != nil {
		newID = fmt.Sprint(rand.Uint64())
	}
	usr = app.NewUser(newID, login)
	c.SetCookie("userID", usr.ID, 3600, "/", "localhost", false, true)
	userSet.AppendUser(*usr)
	c.String(http.StatusOK, "User"+login+" Added id = "+usr.ID)

}

func showUsers(c *gin.Context) {
	cUsers := userSet.GetUsers()
	for _, i := range cUsers {
		e, err := json.Marshal(i)
		if err != nil {
			fmt.Println(err)
		} else {
			c.String(http.StatusOK, string(e))
		}

	}

}
