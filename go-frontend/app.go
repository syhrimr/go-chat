package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var host = "http://skilvul-gc.com"
var backendHost = "http://skilvul-gc.com:7070"
var groupchatBackendHost = "http://skilvul-gc.com:8080"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/register", renderRegisterPage)
	r.GET("/login", renderLoginPage)
	r.GET("/groupchat/list", renderGetJoinedGroupchats)
	r.GET("/groupchat/explore", renderExplorePage)
	r.GET("/groupchat/create", renderCreateRoom)
	r.GET("/profile", renderProfilePage)
	r.GET("/profile/password", renderChangePassword)
	r.GET("/profile/username", renderChangeUsername)
	r.GET("/profile/profpic", renderChangeProfpic)
	r.Static("/css", "./css")
	r.Run(":80")
}

func renderLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title":        "Login",
			"host":         host,
			"account_host": backendHost,
			"gc_host":      groupchatBackendHost,
		},
	)
}

func renderRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title":        "Register",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderGetJoinedGroupchats(c *gin.Context) {
	c.HTML(http.StatusOK, "groupchatlist.html", gin.H{
		"title":        "Groupchats",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderExplorePage(c *gin.Context) {
	c.HTML(http.StatusOK, "explore.html", gin.H{
		"title":        "Explore Group",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderCreateRoom(c *gin.Context) {
	c.HTML(http.StatusOK, "createroom.html", gin.H{
		"title":        "Create a Room",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderProfilePage(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title":        "Profile",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderChangePassword(c *gin.Context) {
	c.HTML(http.StatusOK, "changepassowrd.html", gin.H{
		"title":        "Profile | Change Password",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderChangeUsername(c *gin.Context) {
	c.HTML(http.StatusOK, "changeusername.html", gin.H{
		"title":        "Profile | Change Username",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderChangeProfpic(c *gin.Context) {
	c.HTML(http.StatusOK, "changeprofpic.html", gin.H{
		"title":        "Profile | Change Profile Picture",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}
