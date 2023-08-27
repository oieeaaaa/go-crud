package profile

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

var view = template.Must(template.ParseFiles(
  "views/layout/base.html",
  "views/pages/profile/profile.html",
))

func Render(c *gin.Context) {
  err := view.ExecuteTemplate(c.Writer, "base", nil)

  if err != nil {
    panic(err)
  }
}
