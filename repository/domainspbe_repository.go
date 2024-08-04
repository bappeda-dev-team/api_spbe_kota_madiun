package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type DomainSPBERepository interface {
	Insert(ctx context.Context, tx *sql.Tx, domainspbe domain.DomainSPBE) domain.DomainSPBE
	Update(ctx context.Context, tx *sql.Tx, domainspbe domain.DomainSPBE) domain.DomainSPBE
	Delete(ctx context.Context, tx *sql.Tx, domainspbe domain.DomainSPBE)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.DomainSPBE
	FindById(ctx context.Context, tx *sql.Tx, domainspbeId int) (domain.DomainSPBE, error)
}
