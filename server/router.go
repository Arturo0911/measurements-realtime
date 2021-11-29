package server

import (
	"github.com/Arturo0911/measurements-realtime/controllers"
	. "github.com/Arturo0911/measurements-realtime/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {

	repoLayer := controllers.New()
	router.Group("/realtime").
		GET("/get-reading", repoLayer.GetReading).
		POST("/post-reading", repoLayer.PostReading).
		GET("/get-levels", repoLayer.GetLevels).
		POST("/post-levels", repoLayer.PostLevels)

	router.Use(func(c *gin.Context) {
		c.JSON(404, JsonResponse{
			Success: false,
			Error:   "resource not found!!",
		})
	})

	return router
}
