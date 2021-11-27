package server

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) *gin.Engine {
	router.Group("/realtime").
		GET("get-measures", GetMeasurements).
		GET("get-measure", GetMeasurement).
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
