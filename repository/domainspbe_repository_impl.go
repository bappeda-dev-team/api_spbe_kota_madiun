package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
)

type DomainSPBERepositoryImpl struct {
}

func NewDomainSPBERepository() DomainSPBERepository {
	return &DomainSPBERepositoryImpl{}
}

func (repository *DomainSPBERepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, domainspbe domain.DomainSPBE) domain.DomainSPBE {
	SQL := "INSERT INTO domain_spbe(nama_domain, kode_domain, tahun) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, domainspbe.NamaDomain, domainspbe.KodeDomain, domainspbe.Tahun)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	domainspbe.Id = int(id)
	return domainspbe
}

func (repository *DomainSPBERepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainspbe domain.DomainSPBE) domain.DomainSPBE {
	SQL := "UPDATE domain_spbe SET nama_domain = ?, kode_domain = ?, tahun = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, domainspbe.NamaDomain, domainspbe.KodeDomain, domainspbe.Tahun, domainspbe.Id)
	if err != nil {
		panic(err)
	}
	return domainspbe
}

func (repository *DomainSPBERepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainspbe domain.DomainSPBE) {
	SQL := "DELETE FROM domain_spbe WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, domainspbe.Id)
	if err != nil {
		panic(err)
	}
}

func (repository *DomainSPBERepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.DomainSPBE {
	SQL := "SELECT id, nama_domain, kode_domain, tahun FROM domain_spbe"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var domainspbes []domain.DomainSPBE
	for rows.Next() {
		domainspbe := domain.DomainSPBE{}
		err := rows.Scan(&domainspbe.Id, &domainspbe.NamaDomain, &domainspbe.KodeDomain, &domainspbe.Tahun)
		if err != nil {
			panic(err)
		}
		domainspbes = append(domainspbes, domainspbe)
	}
	return domainspbes
}

func (repository *DomainSPBERepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, domainspbeId int) (domain.DomainSPBE, error) {
	SQL := "SELECT id, nama_domain, kode_domain, tahun FROM domain_spbe WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, domainspbeId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	domainspbe := domain.DomainSPBE{}
	if rows.Next() {
		err := rows.Scan(&domainspbe.Id, &domainspbe.NamaDomain, &domainspbe.KodeDomain, &domainspbe.Tahun)
		if err != nil {
			panic(err)
		}
		return domainspbe, nil
	} else {
		return domainspbe, errors.New("domain spbe tidak ditemukan")
	}
}
