package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Arturo0911/measurements-realtime/connection"
	"github.com/Arturo0911/measurements-realtime/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JsonResponse struct {
	Success bool
	Error   string
}

type ReadingRepo struct {
	Db *gorm.DB
}

type MeasurementsResponse struct {
	IdLevels       int64     `json:"id_niveles"`
	IdReaders      int64     `json:"id_lecturaN"`
	DateWrite      time.Time `json:"fecha_lectura"`
	OxigenMin      float64   `json:"oxigen_min_value"`
	OxigenMax      float64   `json:"oxigen_max_value"`
	TemperatureMin float64   `json:"temperature_min_value"`
	TemperatureMax float64   `json:"temperature_max_value"`
	HumidityMin    float64   `json:"humidity_min_value"`
	HumidityMax    float64   `json:"humidity_max_value"`
	DioxideMin     float64   `json:"dioxide_min_value"`
	DioxideMax     float64   `json:"dioxide_max_value"`
}

func New() *ReadingRepo {
	db := connection.NewInstance()
	db.AutoMigrate(&models.ReadingValues{})
	return &ReadingRepo{
		Db: db,
	}
}

func GetMeasurements(c *gin.Context) {
	c.JSON(200, MeasurementsResponse{
		OxigenMin: 25.22,
	})
}

func (repository *ReadingRepo) GetMeasurements(c *gin.Context) {
	var reading []models.ReadingValues
	err := models.GetLevels(repository.Db, &reading)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err,
			})
		return
	}
	fmt.Println(reading)
	c.JSON(http.StatusOK, reading)
}

func HandleVerification(c *gin.Context) {

	if c.Request.Method == "POST" {
		var measurements MeasurementsResponse
		if err := c.BindJSON(&measurements); err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"oxigen_min_value":      measurements.OxigenMin,
			"oxigen_max_value":      measurements.OxigenMax,
			"temperature_min_value": measurements.TemperatureMin,
			"temperature_max_value": measurements.TemperatureMax,
			"humidity_min_value":    measurements.HumidityMin,
			"humidity_max_value":    measurements.HumidityMax,
			"dioxide_min_value":     measurements.DioxideMin,
			"dioxide_max_value":     measurements.DioxideMax,
		})
		fmt.Println("Parameters are sended!!!")
	}
}
