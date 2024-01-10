package models

import (
	"time"
)

type Unit struct {
	Id          int64 		`gorm:"primaryKey" json:"id"`
	Name        string 		`gorm:"type:varchar(255)" json:"name"`
	Temperature float64 	`gorm:"type:float" json:"temperature"`
	Location    string 		`gorm:"type:varchar(255)" json:"location"`
	Time 		time.Time 	`gorm:"type:datetime" json:"time"`
}