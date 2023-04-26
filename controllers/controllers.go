package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	connectDB "github.com/imagefinder/connect"
	"github.com/imagefinder/models"
)

// Gets image from profile pic the user
func Create(c *gin.Context) {
	var body models.Image

	connectDB.DB.Table("images").AutoMigrate(&body)
}

// Gets image from profile pic the user
func GetImage(c *gin.Context) {
	var image models.Image
	id, exists := c.Get("id")

	formatedID := fmt.Sprintf("%v", id)

	if !exists {
		c.JSON(401, gin.H{
			"message": "Image could not be verified",
		})
		return
	}

	result := connectDB.DB.Table("images").Where("ID = ?", formatedID).First(&image)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "User could not be created",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": image})
}

// Gets image from profile pic the user
func PostImage(c *gin.Context) {
	var body models.Image
	id, exists := c.Get("id")

	formatedID := fmt.Sprintf("%v", id)

	c.BindJSON(&body)

	post := models.Image{
		ID:    formatedID,
		Image: body.Image,
	}
	if !exists {
		c.JSON(401, gin.H{
			"message": "Image could not be verified",
		})
		return
	}
	fmt.Print(id)
	result := connectDB.DB.Table("images").Create(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "User could not be created",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": result})
}
