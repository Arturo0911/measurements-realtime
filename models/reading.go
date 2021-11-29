package models

import (
	"time"

	"gorm.io/gorm"
)

type ReadingValues struct {
	gorm.Model
	Id          int64     `json:"id"`
	DateReading time.Time `json:"date_reading"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Dioxide     float64   `json:"dioxide"`
	Radiation   float64   `json:"radiation"`
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
