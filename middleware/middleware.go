package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	connectDB "github.com/imagefinder/connect"
	jwtbuilder "github.com/imagefinder/helpers"
	"github.com/imagefinder/models"
)

// RequireAuth is a middleware to check if the token is valid
func RequireAuth(c *gin.Context) {
	token, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err = jwtbuilder.VerifyJWTToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var image models.Image
	err = connectDB.DB.Where("ID = ?", token).First(&image).Error
	if err != nil || image.ID == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("image", image)

	c.Next()
}
