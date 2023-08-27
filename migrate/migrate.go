package main

import (
	"github.com/oieeaaaa/go-crud/initializers"
	"github.com/oieeaaaa/go-crud/models"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main() {
  println("migrating stuff...")

  initializers.DB.AutoMigrate(&models.User{})

  println("done!")
}
