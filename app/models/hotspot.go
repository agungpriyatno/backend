package models

import "github.com/agungpriyatno/olap-backend/app/configs/clients"

type Hotspot struct {
	IDLocation   int `json:"-" form:"location"`
	IDConfidence int `json:"-" form:"confidence"`
	IDTime       int `json:"-" form:"time"`
	IDSatelite   int `json:"-" form:"satelite"`
	Total        int `json:"total" form:"total"`

	Location   Location   `json:"location" gorm:"foreignKey:IDLocation"`
	Time       Time       `json:"time" gorm:"foreignKey:IDTime"`
	Confidence Confidence `json:"confidence" gorm:"foreignKey:IDConfidence"`
	Satelite   Satelite   `json:"satelite" gorm:"foreignKey:IDSatelite"`
}

func (data Hotspot) Add() error {
	return clients.DATABASE.Model(&data).Omit("Location", "Time", "Conficence", "Satelite").Create(&data).Error
}
