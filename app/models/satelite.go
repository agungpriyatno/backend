package models

import (
	"github.com/agungpriyatno/olap-backend/app/configs/clients"
)

type Satelite struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
}

func (data Satelite) Add() error {
	return clients.DATABASE.Model(&data).Create(&data).Error
}

// folder models itu untuk model database nya nanti
