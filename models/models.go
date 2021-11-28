package models

import (
	"time"

	"gorm.io/gorm"
)

type Levels struct {
	gorm.Model
	IdLevels       int64     `json:"id_levels"`
	IdReaders      int64     `json:"id_readers"`
	DataWrite      time.Time `json:"data_write"`
	OxigenMin      float64   `json:"oxigen_min_value"`
	OxigenMax      float64   `json:"oxigen_max_value"`
	TemperatureMin float64   `json:"temperature_min_value"`
	TemperatureMax float64   `json:"temperature_max_value"`
	HumidityMin    float64   `json:"humidity_min_value"`
	HumidityMax    float64   `json:"humidity_max_value"`
	DioxideMin     float64   `json:"dioxide_min_value"`
	DioxideMax     float64   `json:"dioxide_max_value"`
}

func CreateLevel(db *gorm.DB, levels *Levels) (err error) {
	err = db.Create(levels).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLevels(db *gorm.DB, levels *[]Levels) (err error) {
	err = db.Find(levels).Error
	if err != nil {
		return err
	}
	return nil
}
