package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/oieeaaaa/go-crud/initializers"
	"github.com/oieeaaaa/go-crud/models"
)

func GetThoughts(c *gin.Context) {
  var thoughts []models.Thoughts

  initializers.DB.Find(&thoughts)

  c.HTML(200, "thoughts", gin.H{
    "Thoughts": thoughts,
  });
}

func CreateThought(c *gin.Context) {
  var thought models.Thoughts
  thought.Say = c.PostForm("Say")

  if thought.Say == "" {
    c.JSON(400, gin.H{"error": "Say is required"})
    return
  }

  // create thought
  initializers.DB.Create(&thought)

  var thoughts []models.Thoughts

  initializers.DB.Find(&thoughts)

  c.HTML(200, "thoughts", gin.H{
    "Thoughts": thoughts,
  });
}

func DeleteThought(c *gin.Context) {
  var thought models.Thoughts
  id := c.Param("id")

  initializers.DB.Where("id = ?", id).Delete(&thought)

  var thoughts []models.Thoughts

  initializers.DB.Find(&thoughts)
}
