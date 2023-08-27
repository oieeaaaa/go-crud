package main

import (
	"github.com/oieeaaaa/go-crud/controllers"
	"github.com/oieeaaaa/go-crud/initializers"
	"github.com/oieeaaaa/go-crud/views/pages/about"
	"github.com/oieeaaaa/go-crud/views/pages/home"
	"github.com/oieeaaaa/go-crud/views/pages/profile"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main() {
  r := initializers.Router()

  // ==================== API ====================

  r.POST("/users/search", controllers.SearchUser)
  /* r.POST("/user", controllers.CreateUser)
  r.PUT("/user/:id", controllers.UpdateUser)
  r.GET("/user/:id", controllers.GetUser)
  r.GET("/users", controllers.GetUsers)
  r.DELETE("/user/:id", controllers.DeleteUser) */

  // ==================== HTML ====================

  r.GET("/", home.Render)
  r.GET("/about", about.Render)
  r.GET("/profile", profile.Render)

  // ==================== STATIC ====================

  r.Static("/assets", "./assets")

  r.Run() // listen and serve on 0.0.0.0:8080
}
