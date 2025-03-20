package main

import (
	"bookstore_project/controllers"
	"bookstore_project/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts",controllers.PostsIndex)
	router.GET("/posts/:id",controllers.PostShow)
	router.PUT("/posts/:id",controllers.PostUpdate)
	router.DELETE("/posts/:id",controllers.PostDelete)
	router.Run() // listen and serve on 0.0.0.0:8080

	
}