package main

import (
	"html/template"

	"github.com/oieeaaaa/go-crud/controllers"
	"github.com/oieeaaaa/go-crud/initializers"
	"github.com/oieeaaaa/go-crud/views/pages/about"
	"github.com/oieeaaaa/go-crud/views/pages/home"
	"github.com/oieeaaaa/go-crud/views/pages/profile"
	"github.com/oieeaaaa/go-crud/views/pages/user"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main() {
  r := initializers.Router()

  // ==================== API ====================

  r.POST("/users/search", controllers.SearchUser)
  r.POST("/thoughts", controllers.CreateThought)
  /* r.POST("/user", controllers.CreateUser)
  r.PUT("/user/:id", controllers.UpdateUser)
  r.GET("/user/:id", controllers.GetUser)
  r.GET("/users", controllers.GetUsers)
  r.DELETE("/user/:id", controllers.DeleteUser) */

  // ==================== HTML ====================

  html := template.Must(template.ParseGlob("views/partials/*.html"))

  r.SetHTMLTemplate(html)

  r.GET("/", home.Render)
  r.GET("/about", about.Render)
  r.GET("/profile", profile.Render)
  r.GET("/user", user.Render)

  // ==================== STATIC ====================

  r.Static("/assets", "./assets")

  r.Run() // listen and serve on 0.0.0.0:8080
}
