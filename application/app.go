package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/bookstore-api/controllers"
	"github.com/pt-abhishek/bookstore-api/databases/elasticsearch"
	"github.com/pt-abhishek/bookstore-api/databases/mysql"
)

//StartApplication starts the application
func StartApplication() {

	//init connections to the database
	elasticsearch.ElasticClient.Init()
	mysql.SQLClient.Init()

	//init go router
	router := gin.Default()

	router.GET("/books/search", controllers.SearchBooks)
	router.GET("/books", controllers.FindAllWithPagination)

	router.Run(":8082")
}
