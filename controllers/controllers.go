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

type LevelsRepo struct {
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

func New() *LevelsRepo {
	db := connection.NewInstance()
	db.AutoMigrate(&models.Niveles{})
	return &LevelsRepo{
		Db: db,
	}
}

// func GetMeasurements(c *gin.Context) {
// 	c.JSON(200, MeasurementsResponse{
// 		OxigenMin: 25.22,
// 	})
// }

func (repository *LevelsRepo) GetMeasurements(c *gin.Context) {
	var niveles []models.Niveles
	err := models.GetLevels(repository.Db, &niveles)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err,
			})
		return
	}
	fmt.Println(niveles)
	c.JSON(http.StatusOK, niveles)
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
