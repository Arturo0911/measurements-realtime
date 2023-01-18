package models

import (
	"time"

	"gorm.io/gorm"
)

type ReadingValues struct {
	//gorm.Model
	Id          int64     `json:"id"`
	DateReading time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	PH          float64   `json:"ph"`
}

func CreateReading(db *gorm.DB, reading *ReadingValues) (err error) {
	err = db.Create(reading).Error
	if err != nil {
		return err
	}
	return nil
}

func GetReadingValues(db *gorm.DB, reading *[]ReadingValues) (err error) {
	err = db.Find(reading).Error
	if err != nil {
		return err
	}
	return nil
}
