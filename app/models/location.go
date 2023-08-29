package models

import (
	"github.com/agungpriyatno/olap-backend/app/configs/clients"
)

type Location struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Pulau     string `json:"pulau" form:"pulau"`
	Provinsi  string `json:"provinsi" form:"provinsi"`
	Kota      string `json:"kota" form:"kota"`
	Kecamatan string `json:"kecamatan" form:"kecamatan"`
	Desa      string `json:"desa" form:"desa"`
}

func (data Location) Add() error {
	return clients.DATABASE.Model(&data).Create(&data).Error
}
