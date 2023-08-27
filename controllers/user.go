package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/oieeaaaa/go-crud/initializers"
	"github.com/oieeaaaa/go-crud/models"
)

func CreateUser(c *gin.Context) {
  // get data off req body

  var userReq struct {
    Name string
    Email string
  }

  c.Bind(&userReq);

  // create user
  user := models.User{
    Name: userReq.Name,
    Email: userReq.Email,
  }

  result := initializers.DB.Create(&user)

  if result.Error != nil {
    c.JSON(500, gin.H{
      "message": "error creating user",
    })
    return
  }

  c.JSON(200, gin.H{
    "message": "user created",
    "user": user,
  })
}

func GetUsers(c *gin.Context) {
  var users []models.User

  initializers.DB.Find(&users)

  c.JSON(200, gin.H{
    "users": users,
  })
}

func GetUser(c *gin.Context) {
  // Get id from params
  id := c.Param("id")

  var user models.User

  initializers.DB.First(&user, id)

  c.HTML(200, "users", user)
}

func UpdateUser(c *gin.Context) {
  // Get id from params
  id := c.Param("id")

  var userReq struct {
    Name string
    Email string
  }

  c.Bind(&userReq);

  var user models.User

  initializers.DB.First(&user, id)

  user.Name = userReq.Name
  user.Email = userReq.Email

  initializers.DB.Model(&user).Updates(user)

  c.JSON(200, gin.H{
    "message": "user updated",
    "user": user,
  })
}

func DeleteUser(c *gin.Context) {
  // Get id from params
  id := c.Param("id")

  initializers.DB.Delete(&models.User{}, id)

  c.JSON(200, gin.H{
    "message": "user deleted",
  })
}

func SearchUser(c *gin.Context) {
  // Get query from FormData
  search := c.PostForm("search")

  var Users []models.User

  initializers.DB.Where("name ILIKE ?", "%" + search + "%").Find(&Users)

  if len(Users) == 0 {
    c.HTML(200, "usersNotFound", nil)
    return
  }

  c.HTML(200, "userSearchResults", gin.H{
    "users": Users,
  })
}
