package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/YuraSich/ShelterGame/app"

	"github.com/gin-gonic/gin"
)

var (
	users []*app.User
)

func main() {
	r := gin.Default()
	users = make([]*app.User, 0)
	_, err := app.NewAPIServer("test_db")
	if err != nil {
		log.Fatal(err)
	}
	r.GET("/login", loginHandlerGET)
	r.POST("/login", loginHandlerPOST)
	r.GET("/registration", regHandlerGET)
	r.POST("/registration", regHandlerPOST)
	r.GET("/", indexHandler)
	r.Run(":3030")
}

func indexHandler(c *gin.Context) {
	auth, err := c.Cookie("auth")
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	if auth != "yes" {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	t, err := template.ParseFiles("view/static/mainPage.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "HUYAMBA-loginHandlerGET")
	}
	t.Execute(c.Writer, c.Request)
}

func loginHandlerGET(c *gin.Context) {
	t, err := template.ParseFiles("view/static/login.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "HUYAMBA-loginHandlerGET")
	}
	t.Execute(c.Writer, c.Request)
}

func loginHandlerPOST(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println("User = " + username + "\tpassword = " + password)
	for _, i := range users {
		if i.Login == username {
			log.Println("Login correct")
			if i.Password == password {
				log.Println("Password correct")
				c.SetCookie("auth", "yes", 3600, "/", "localhost", false, true)
				c.Redirect(http.StatusFound, "/")
				c.Abort()
			}
		}
	}
	//loginHandlerGET(c)
}

func regHandlerGET(c *gin.Context) {
	t, err := template.ParseFiles("view/static/registration.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "HUYAMBA-regHandlerGET")
	}
	t.Execute(c.Writer, c.Request)
}

func regHandlerPOST(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println("User = " + username + "\tpassword = " + password)
	usr := app.NewUser(
		"",
		username,
		"",
		password,
		"",
	)
	users = append(users, usr)
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
}
