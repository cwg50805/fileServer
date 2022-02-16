package routes

import (
	"CreativeServer/controllers"
	v1 "CreativeServer/controllers/v1"
	_ "CreativeServer/docs"
	"CreativeServer/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if mode := gin.Mode(); mode == gin.DebugMode {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.GET("/heartBeat", controllers.HeartBeat)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(middleware.MiddleWare)
	{
		apiv1.GET("/file/*localSystemFilePath", v1.GetFile)
		apiv1.POST("/file/*localSystemFilePath", v1.CreateFile)
		apiv1.PATCH("/file/*localSystemFilePath", v1.UpdateFile)
		apiv1.DELETE("/file/*localSystemFilePath", v1.DeleteFile)
	}

	return router
}
