package main

import (
	"gee"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gee.New()
	html1, _ := ioutil.ReadFile("C:/Users/Administrator/Desktop/1.html")
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, string(html1))
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello1", func(c *gee.Context) {
		// c.JSON(http.StatusOK, c.Query("name"))
		c.JSON(http.StatusInternalServerError, c.Query("name"))
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
