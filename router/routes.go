package router

import (
	"API_GO/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	//inicializar handler
	handlers.InitHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.GetUsers)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.DELETE("/opening", handlers.DeleteOpeningHandler)
		v1.PUT("/opening", handlers.UpdateOpeningHandler)
		v1.GET("/openingALL", handlers.GetAllOpenningHandler)
	}
}
