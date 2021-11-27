package server

import "github.com/gin-gonic/gin"

func Server(router *gin.Engine) *gin.Engine {
	router.Group("/").
		GET("get-measure", GetMeasurements).
		GET("send-measure", SendMeasurements).
		POST("post-measure", HandleVerification)
		// possible POST method to send Mysql

	router.Use(func(c *gin.Context) {
		c.JSON(404, JsonResponse{
			Success: false,
			Error:   "resource not found!!",
		})
	})

	return router
}
