package router

import (
	controllers "C2-simple-rest-gin/controller"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBook)
	router.GET("/books/:bookID", controllers.GetBook)
	router.POST("/books", controllers.AddBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
