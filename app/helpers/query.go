package helpers

import (
	"github.com/agungpriyatno/olap-backend/app/configs/clients"
	"github.com/agungpriyatno/olap-backend/app/models"
	"github.com/agungpriyatno/olap-backend/app/models/payload"
	"github.com/xdbsoft/olap"
)

func SQLtoCube() (olap.Cube, error) {
	var count int64
	var list []models.Hotspot
	var data [][]interface{}
	// var dataLocation [][]interface{}

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
			item.Location.Pulau,
			item.Time.Tahun,
			item.Confidence.Level,
			item.Satelite.Name,
			item.Total,
		})
	}

	cube.AddRows([]string{"location", "time", "confidence", "satelite", "total"}, data)

	return cube, nil
}

// kita permudah biar lebih enak dilihat
// Asu laa tidak bisa

func CubeLocation(param payload.Location) (olap.Cube, error) {
	var count int64
	var list []models.Hotspot
	var data [][]interface{}

	cube := olap.Cube{
		Dimensions: []string{"location", "time", "confidence", "satelite"},
		Fields:     []string{"total"},
	}

	where := make(map[string]interface{})

	if param.Kecamatan != "" {
		where["Location.Kecamatan"] = param.Kecamatan
	}

	if param.Kota != "" {
		where["Location.Kota"] = param.Kota
	}

	if param.Provinsi != "" {
		where["Location.Provinsi"] = param.Provinsi
	}

	if param.Pulau != "" {
		where["Location.Pulau"] = param.Pulau
	}

	print(where["Location.Kota"])

	if err := clients.DATABASE.Model(&models.Hotspot{}).
		Joins("Location").Preload("Confidence").
		Preload("Time").Preload("Satelite").Where(where).
		Find(&list).Count(&count).Error; err != nil {
		return cube, err
	}

	for i := 0; i < int(count); i++ {
		var value string
		item := list[i]
		if param.Kecamatan != "" {
			value = item.Location.Desa
		} else if param.Kota != "" {
			value = item.Location.Kecamatan
		} else if param.Provinsi != "" {
			value = item.Location.Kota
		} else if param.Pulau != "" {
			value = item.Location.Provinsi
		}
		data = append(data, []interface{}{
			value,
			item.Time.Tahun,
			item.Confidence.Level,
			item.Satelite.Name,
			item.Total,
		})
	}

	cube.AddRows([]string{"location", "time", "confidence", "satelite", "total"}, data)

	return cube, nil
}

func CubeTime(param payload.Time) (olap.Cube, error) {
	var count int64
	var list []models.Hotspot
	var data [][]interface{}

	cube := olap.Cube{
		Dimensions: []string{"location", "time", "confidence", "satelite"},
		Fields:     []string{"total"},
	}

	where := make(map[string]interface{})

	if param.Bulan != "" {
		where["Time.Bulan"] = param.Bulan
	}

	if param.Kuartal != "" {
		where["Time.Kuartal"] = param.Kuartal
	}

	if param.Semester != "" {
		where["Time.Semester"] = param.Semester
	}

	if param.Tahun != "" {
		where["Time.Tahun"] = param.Tahun
	}

	if err := clients.DATABASE.Model(&models.Hotspot{}).
		Preload("Location").Preload("Confidence").
		Joins("Time").Preload("Satelite").Where(where).
		Find(&list).Count(&count).Error; err != nil {
		return cube, err
	}

	for i := 0; i < int(count); i++ {
		var value string
		item := list[i]
		if param.Bulan != "" {
			value = item.Time.Hari
		} else if param.Kuartal != "" {
			value = item.Time.Bulan
		} else if param.Semester != "" {
			value = item.Time.Kuartal
		} else if param.Tahun != "" {
			value = item.Time.Semester
		}
		data = append(data, []interface{}{
			item.Location.Pulau,
			value,
			item.Confidence.Level,
			item.Satelite.Name,
			item.Total,
		})
	}

	cube.AddRows([]string{"location", "time", "confidence", "satelite", "total"}, data)

	return cube, nil
}

func CubeTimeLocation(param payload.Time, loc payload.Location) (olap.Cube, error) {
	var count int64
	var list []models.Hotspot
	var data [][]interface{}

	cube := olap.Cube{
		Dimensions: []string{"location", "time", "confidence", "satelite"},
		Fields:     []string{"total"},
	}

	where := make(map[string]interface{})
	whereLoc := make(map[string]interface{})

	if param.Bulan != "" {
		where["Time.Bulan"] = param.Bulan
	}

	if param.Kuartal != "" {
		where["Time.Kuartal"] = param.Kuartal
	}

	if param.Semester != "" {
		where["Time.Semester"] = param.Semester
	}

	if param.Tahun != "" {
		where["Time.Tahun"] = param.Tahun
	}

	if loc.Kecamatan != "" {
		whereLoc["Location.Kecamatan"] = loc.Kecamatan
	}

	if loc.Kota != "" {
		whereLoc["Location.Kota"] = loc.Kota
	}

	if loc.Provinsi != "" {
		whereLoc["Location.Provinsi"] = loc.Provinsi
	}

	if loc.Pulau != "" {
		whereLoc["Location.Pulau"] = loc.Pulau
	}

	if err := clients.DATABASE.Model(&models.Hotspot{}).
		Joins("Location").Preload("Confidence").
		Joins("Time").Preload("Satelite").Where(where).Where(whereLoc).
		Find(&list).Count(&count).Error; err != nil {
		return cube, err
	}

	for i := 0; i < int(count); i++ {
		var time string
		var location string

		item := list[i]

		if param.Bulan != "" {
			time = item.Time.Hari
		} else if param.Kuartal != "" {
			time = item.Time.Bulan
		} else if param.Semester != "" {
			time = item.Time.Kuartal
		} else if param.Tahun != "" {
			time = item.Time.Semester
		} else {
			time = item.Time.Tahun
		}

		if loc.Kecamatan != "" {
			location = item.Location.Desa
		} else if loc.Kota != "" {
			location = item.Location.Kecamatan
		} else if loc.Provinsi != "" {
			location = item.Location.Kota
		} else if loc.Pulau != "" {
			location = item.Location.Provinsi
		} else {
			location = item.Location.Pulau
		}
		data = append(data, []interface{}{
			location,
			time,
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
