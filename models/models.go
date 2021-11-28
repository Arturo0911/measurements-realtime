package models

import (
	"time"

	"gorm.io/gorm"
)

type Levels struct {
	gorm.Model
	IdLevels       int64
	IdReaders      int64
	DataWrite      time.Time
	OxigenMin      float64
	OxigenMax      float64
	TemperatureMin float64
	TemperatureMax float64
	HumidityMin    float64
	HumidityMax    float64
	DioxideMin     float64
	DioxideMax     float64
}

func CreateLevel(db *gorm.DB, levels *Levels) (err error) {
	err = db.Create(levels).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLevels(db *gorm.DB, levels *Levels) (err error) {
	err = db.Find(levels).Error
	if err != nil {
		return err
	}
	return nil
}
