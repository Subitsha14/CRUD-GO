package controllers

import (
	"CRUD-GO/initializers"
	"CRUD-GO/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get the post data
	var data struct {
		Title string
		Body  string
	}
	c.Bind(&data)
	//creating the posts
	post := models.Post{Title: data.Title, Body: data.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsShowAll(c *gin.Context) {
	//fetching all the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShowById(c *gin.Context) {
	//fetching single post by ID

	id := c.Param("id")
	var post models.Post
	initializers.DB.Find(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get id then get the data
	// find the post using id then update it

	id := c.Param("id")
	var data struct {
		Title string
		Body  string
	}
	c.Bind(&data)
	var post models.Post
	initializers.DB.Find(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: data.Title,
		Body:  data.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//get the id then delete the post
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)

}
