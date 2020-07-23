package main

import (
	"log"
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
	r.POST("/addUser", addUser)
	r.Run()
}

func addUser(c *gin.Context) {
	id := c.Query("id")
	log.Print("ID = " + id)
	login := c.Query("login")
	log.Println("login = " + login)
	if id == "" || login == "" {
		c.String(http.StatusBadRequest, "ID Or Login is invalid")
		return
	}
	c.String(http.StatusOK, "User ")
	userSet.Append(id, login)
}

func showUsers(c *gin.Context) {
	cUsers := userSet.GetUsers()
	for i := 0; i < len(cUsers); i++ {
		c.JSON(200, gin.H{
			cUsers[i].ID: cUsers[i].Login,
		})
	}

}
