package models

import (
	"github.com/agungpriyatno/olap-backend/app/configs/clients"
)

type Confidence struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Level string `json:"level" form:"level"`
}

func (data Confidence) Add() error {
	return clients.DATABASE.Model(&data).Create(&data).Error
}
