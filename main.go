package main

import (
	"CRUD-GO/controllers"
	"CRUD-GO/initializers"
	"CRUD-GO/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.PUT("/post/:id", controllers.PostsUpdate)
	r.GET("/post", controllers.PostsShowAll)
	r.GET("/post/:id", controllers.PostsShowById)
	r.DELETE("/post/:id", controllers.PostsDelete)
	r.Run()
}
