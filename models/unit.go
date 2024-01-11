package models

import (
	"database/sql/driver"
	"time"
)

type Unit struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Temperature float64   `gorm:"type:float" json:"temperature"`
	Location    string    `gorm:"type:varchar(255)" json:"location"`
	Time        CustomTime `gorm:"type:datetime" json:"time"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) Scan(value interface{}) error {
	tt, err := time.Parse("2006-01-02 15:04:05", string(value.([]uint8)))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t CustomTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time.Format("2006-01-02 15:04:05"), nil
}