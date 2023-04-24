package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imagefinder/models"
)

// Gets image from profile pic the user
func GetImage(c *gin.Context) {
	image, exists := c.Get("image")

	if !exists {
		c.JSON(401, gin.H{
			"message": "Image could not be verified",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": image})
}


// Gets image from profile pic the user
func PostImage(c *gin.Context) {
	var body models.Image
	image, exists := c.Get("image")

	if !exists {
		c.JSON(401, gin.H{
			"message": "Image could not be verified",
		})
		return
	}
	c.BindJSON(&body)

	post := models.Image{
		ID:       body.ID,
		Image:    body.Image,
	}
	result := connectDB.DB.Create(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "User could not be created",
		})
		return
	
	
	



	c.JSON(200, gin.H{
		"message": image})
}

