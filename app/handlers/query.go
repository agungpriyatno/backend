package handlers

import (
	"net/http"

	"github.com/agungpriyatno/olap-backend/app/helpers"
	"github.com/gofiber/fiber/v2"
)

func Query(ctx *fiber.Ctx) error {
	dimension := ctx.Params("dimension")
	point := ctx.Query("pt")
	dim := ctx.Query("dim")

	cube, err := helpers.SQLtoCube()
	if err != nil {
		return ResponseJson(ctx, http.StatusBadGateway, err.Error())
	}

	if point != "" && dim != "" {
		cube = cube.Slice(dimension, point)
		cube = cube.RollUp([]string{dim}, cube.Fields, helpers.Sum, []interface{}{0})
	} else {
		cube = cube.RollUp([]string{dimension}, cube.Fields, helpers.Sum, []interface{}{0})
	}

	return ResponseJson(ctx, http.StatusOK, cube.Rows())
}
