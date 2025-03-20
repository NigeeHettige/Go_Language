package controllers

import (
	"bookstore_project/initializers"
	"bookstore_project/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data of request body

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {

		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {
	//Get id of url
	id := c.Param("id")
	var posts models.Post
	initializers.DB.First(&posts, id)

	c.JSON(200, gin.H{
		"post": posts,
	})

}

func PostUpdate(c *gin.Context) {
	//Get the id of the url
	id := c.Param("id")

	//Get the data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	//Update it

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	//Respond it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context){
	//Get the id of url
	id := c.Param("id")

	//Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}