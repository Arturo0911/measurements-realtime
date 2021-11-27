package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Success bool
	Error   string
}

type MeasurementsResponse struct {
	Oxigen      float64 `json:"oxigen_value"`
	Temperature float64 `json:"temperature_value"`
	Humidity    float64 `json:"humidity_value"`
	Dioxide     float64 `json:"dioxide_value"`
}

func GetMeasurements(c *gin.Context) {
	c.JSON(200, MeasurementsResponse{
		Oxigen: 25.22,
	})
}

func SendMeasurements(c *gin.Context) {
	c.JSON(200, MeasurementsResponse{})
}

func HandleVerification(c *gin.Context) {

	if c.Request.Method == "POST" {
		var measurements MeasurementsResponse
		if err := c.BindJSON(&measurements); err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"oxigen_value":      measurements.Oxigen,
			"temperature_value": measurements.Temperature,
			"humidity_value":    measurements.Humidity,
			"dioxide_value":     measurements.Dioxide,
		})
		fmt.Println("Parameters are sended!!!")
	}
}
