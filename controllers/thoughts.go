package controllers

import (
	"strconv"

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

  res := initializers.DB.Create(&thought)

  if res.Error != nil {
    c.JSON(400, gin.H{"error": res.Error})
    return
  }

  c.HTML(200, "thought", thought);
}

func EditThought(c *gin.Context) {
  var thought models.Thoughts

  // client requests
  id := c.Param("id")
  thought.Say = c.PostForm("OldSay")

  // important: convert id back
  newId, _ := strconv.ParseUint(id, 0, 32)
  thought.ID = uint(newId)

  // slap that update
  initializers.DB.Model(&thought).Where("id = ?", id).Update("Say", thought.Say)
  initializers.DB.Find(&thought, id)

  c.HTML(200, "thought", thought)
}

func DeleteThought(c *gin.Context) {
  var thought models.Thoughts
  id := c.Param("id")

  initializers.DB.Where("id = ?", id).Delete(&thought)

  var thoughts []models.Thoughts

  initializers.DB.Find(&thoughts)
}
