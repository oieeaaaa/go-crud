package about

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

var view = template.Must(template.ParseFiles(
  "views/layout/base.html",
  "views/pages/about/about.html",

  "views/icons.html",
  "views/partials/_modal.html",
))

func Render(c *gin.Context) {
  err := view.ExecuteTemplate(c.Writer, "base", nil)

  if err != nil {
    panic(err)
  }
}
