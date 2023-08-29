package helpers

import (
	"github.com/agungpriyatno/olap-backend/app/models/payload"
	"github.com/gofiber/fiber/v2"
)

func GetPayload(ctx *fiber.Ctx) (payload.Location, payload.Time) {
	var location payload.Location
	var time payload.Time

	location.Pulau = ctx.Query("pulau", "")
	location.Provinsi = ctx.Query("provinsi", "")
	location.Kota = ctx.Query("kota", "")
	location.Kecamatan = ctx.Query("kecamatan", "")

	time.Tahun = ctx.Query("tahun", "")
	time.Semester = ctx.Query("semester", "")
	time.Kuartal = ctx.Query("kuartal", "")
	time.Bulan = ctx.Query("bulan", "")

	return location, time

}
