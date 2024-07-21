package helper

import (
	"api_spbe_kota_madiun/model/web"
	"database/sql"
)

func ConvertNullInt32ToProsBisPohonKinerja(id sql.NullInt32) *web.ProsBisPohonKinerjaRespons {
	if id.Valid {
		return &web.ProsBisPohonKinerjaRespons{
			ID: int(id.Int32),
		}
	}
	return nil
}

func ConvertNullInt32ToPointer(nullInt sql.NullInt32) *int {
	if nullInt.Valid {
		val := int(nullInt.Int32)
		return &val
	}
	return nil
}
