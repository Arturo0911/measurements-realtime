package models

import (
	"time"

	"gorm.io/gorm"
)

// type Niveles struct {
// 	gorm.Model
// 	IdLevels       int64     `json:"id_niveles"`
// 	IdReaders      int64     `json:"id_lecturaN"`
// 	DateWrite      time.Time `json:"fecha_lectura"`
// 	OxigenMin      float64   `json:"uvMin"`
// 	OxigenMax      float64   `json:"uvMax"`
// 	TemperatureMin float64   `json:"temperaturaMin"`
// 	TemperatureMax float64   `json:"temperaturaMax"`
// 	HumidityMin    float64   `json:"humedadMin"`
// 	HumidityMax    float64   `json:"humedadMax"`
// 	DioxideMin     float64   `json:"co2Min"`
// 	DioxideMax     float64   `json:"co2Max"`
// }

type ReadingValues struct {
	gorm.Model
	Id          int64     `json:"id"`
	DateReading time.Time `json:"date_reading"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Dioxide     float64   `json:"dioxide"`
	Radiation   float64   `json:"radiation"`
}

func CreateLevel(db *gorm.DB, reading *ReadingValues) (err error) {
	err = db.Create(reading).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLevels(db *gorm.DB, reading *[]ReadingValues) (err error) {
	err = db.Find(reading).Error
	if err != nil {
		return err
	}
	return nil
}
