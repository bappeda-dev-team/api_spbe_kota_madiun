package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type ReferensiArsitekturRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur) domain.ReferensiArsitektur
	Update(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur) domain.ReferensiArsitektur
	Delete(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ReferensiArsitektur
	FindByKodeRef(ctx context.Context, tx *sql.Tx, kodeReferensi string) *helper.TreeNode
	FindById(ctx context.Context, tx *sql.Tx, referensiarsitekturId int) (domain.ReferensiArsitektur, error)
}
