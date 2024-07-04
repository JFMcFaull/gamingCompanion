package main

import (
	"github.com/JFMcFaull/gamingCompanion/controllers"
	"github.com/JFMcFaull/gamingCompanion/initalizers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"} // Adjust this to your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	r.POST("/games", controllers.GamesCreate)
	r.PUT("/games/:id", controllers.GamesUpdate)
	r.GET("/games", controllers.GamesIndex)
	r.GET("/games/:id", controllers.GamesShow)
	r.DELETE("/games/:id", controllers.GamesDelete)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (or any desired port)
}
