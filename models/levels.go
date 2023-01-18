package models

import (
	"time"

	"gorm.io/gorm"
)

// soil levels
type LevelsValues struct {
	//gorm.Model
	Id          int64     `json:"id"`
	DateReading time.Time `gorm:"type:timestamp;default:current_timestamp" json:"date_reading"`
	PH          float64   `json:"ph_value"`
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
