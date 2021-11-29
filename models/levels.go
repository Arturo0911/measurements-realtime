package models

import (
	"time"

	"gorm.io/gorm"
)

type LevelsValues struct {
	gorm.Model
	Id             int64     `json:"id"`
	DateReading    time.Time `json:"date_reading"`
	TemperatureMin float64   `json:"temperature_min"`
	TemperatureMax float64   `json:"temperature_max"`
	HumidityMin    float64   `json:"humidity_min"`
	HumidityMax    float64   `json:"humidity_max"`
	DioxideMin     float64   `json:"dioxide_min"`
	DioxideMax     float64   `json:"dioxide_max"`
	UvMin          float64   `json:"uv_min"`
	UvMax          float64   `json:"uv_max"`
}

func CreateLevels(db *gorm.DB, levels *LevelsValues) (err error) {
	err = db.Create(levels).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLevelsValues(db *gorm.DB, levels *[]LevelsValues) (err error) {
	err = db.Find(levels).Error
	if err != nil {
		return err
	}
	return nil
}
