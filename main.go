package main

import (
	"html/template"

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

  // ==================== API ====================

  // routes
  r.POST("/users/search", controllers.SearchUser)
  /* r.POST("/user", controllers.CreateUser)
  r.PUT("/user/:id", controllers.UpdateUser)
  r.GET("/user/:id", controllers.GetUser)
  r.GET("/users", controllers.GetUsers)
  r.DELETE("/user/:id", controllers.DeleteUser) */

  // ==================== API ====================

  // ==================== HTML ====================

  // serve html
	r.LoadHTMLGlob("templates/*")

  // Html nested templates
  tmpl := make(map[string]*template.Template)
  tmpl["about.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/about.html"))
  tmpl["profile.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/profile.html"))

  // Html routes
  r.GET("/about", func(c *gin.Context) {
    tmpl["about.html"].ExecuteTemplate(c.Writer, "base", nil)
  })

  r.GET("/profile", func(c *gin.Context) {
    tmpl["profile.html"].ExecuteTemplate(c.Writer, "base", nil)
  })

  r.Run() // listen and serve on 0.0.0.0:8080

  // ==================== HTML ====================
}
