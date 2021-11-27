package server

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Success bool
	Error   string
}

type measurementsResponse struct {
	Oxigen      float64 `json:"oxigen_value"`
	Temperature float64 `json:"temperature_value"`
	Humidity    float64 `json:"humidity_value"`
	Dioxide     float64 `json:"dioxide_value"`
}

func GetMeasurements(c *gin.Context) {
	c.JSON(200, measurementsResponse{
		Oxigen: 25.22,
	})
}

func SendMeasurements(c *gin.Context) {
	c.JSON(200, measurementsResponse{})
}
