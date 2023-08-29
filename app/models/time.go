package models

import (
	"github.com/agungpriyatno/olap-backend/app/configs/clients"
)

type Time struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Tahun    string `json:"tahun" form:"tahun"`
	Semester string `json:"semester" form:"semester"`
	Kuartal  string `json:"kuartal" form:"kuartal"`
	Bulan    string `json:"bulan" form:"bulan"`
	Hari     string `json:"hari" form:"hari"`
}

func (data Time) Add() error {
	return clients.DATABASE.Model(&data).Create(&data).Error
}
