package models

import (
	"time"

	"gorm.io/gorm"
)

type Niveles struct {
	gorm.Model
	IdLevels       int64     `json:"id_niveles"`
	IdReaders      int64     `json:"id_lecturaN"`
	DateWrite      time.Time `json:"fecha_lectura"`
	OxigenMin      float64   `json:"uvMin"`
	OxigenMax      float64   `json:"uvMax"`
	TemperatureMin float64   `json:"temperaturaMin"`
	TemperatureMax float64   `json:"temperaturaMax"`
	HumidityMin    float64   `json:"humedadMin"`
	HumidityMax    float64   `json:"humedadMax"`
	DioxideMin     float64   `json:"co2Min"`
	DioxideMax     float64   `json:"co2Max"`
}

func CreateLevel(db *gorm.DB, niveles *Niveles) (err error) {
	err = db.Create(niveles).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLevels(db *gorm.DB, niveles *[]Niveles) (err error) {
	err = db.Find(niveles).Error
	if err != nil {
		return err
	}
	return nil
}
