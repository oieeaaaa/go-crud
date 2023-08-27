package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oieeaaaa/go-crud/controllers"
	"github.com/oieeaaaa/go-crud/initializers"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main() {
  r := gin.Default()

  // routes
  /* r.POST("/user", controllers.CreateUser)
  r.PUT("/user/:id", controllers.UpdateUser)
  r.GET("/user/:id", controllers.GetUser)
  r.GET("/users", controllers.GetUsers)
  r.DELETE("/user/:id", controllers.DeleteUser) */

  r.POST("/users/search", controllers.SearchUser)

  // serve html
	r.LoadHTMLGlob("templates/*")

  r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

  r.Run() // listen and serve on 0.0.0.0:8080
}
