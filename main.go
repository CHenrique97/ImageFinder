package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	connectDB "github.com/imagefinder/connect"
	"github.com/imagefinder/initializers"
)

func init() {
	initializers.LoadEnv()
	connectDB.InitConnector()
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your client's URL
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/getImage", controllers.GetImage, controllers.Validate)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/postImage", controllers.PostImage, controllers.Validate)
	r.Run(":" + os.Getenv("PORT"))

}
