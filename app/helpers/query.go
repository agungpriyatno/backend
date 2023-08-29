package helpers

import (
	"github.com/agungpriyatno/olap-backend/app/configs/clients"
	"github.com/agungpriyatno/olap-backend/app/models"
	"github.com/xdbsoft/olap"
)

func SQLtoCube() (olap.Cube, error) {
	var count int64
	var list []models.Hotspot
	var data [][]interface{}
	var dataLocation [][]interface{}

	cube := olap.Cube{
		Dimensions: []string{"location", "time", "confidence", "satelite"},
		Fields:     []string{"total"},
	}

	if err := clients.DATABASE.Model(&models.Hotspot{}).
		Preload("Location").Preload("Confidence").
		Preload("Time").Preload("Satelite").
		Find(&list).Count(&count).Error; err != nil {
		return cube, err
	}

	for i := 0; i < int(count); i++ {
		item := list[i]
		data = append(data, []interface{}{
			dataLocation,
			item.Time.Tahun,
			item.Confidence.Level,
			item.Satelite.Name,
			item.Total,
		})
	}

	cube.AddRows([]string{"location", "time", "confidence", "satelite", "total"}, data)

	return cube, nil
}

func Sum(aggregate, value []interface{}) []interface{} {
	s := aggregate[0].(int)
	s += value[0].(int)
	return []interface{}{s}
}
