package handlers

import (
	"net/http"

	"github.com/agungpriyatno/olap-backend/app/models"
	"github.com/gofiber/fiber/v2"
)

func AddHotspot(ctx *fiber.Ctx) error {
	var data models.Hotspot

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddLocation(ctx *fiber.Ctx) error {
	var data models.Location

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddTime(ctx *fiber.Ctx) error {
	var data models.Time

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddConfidence(ctx *fiber.Ctx) error {
	var data models.Confidence

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddSatelite(ctx *fiber.Ctx) error {
	var data models.Satelite

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}
