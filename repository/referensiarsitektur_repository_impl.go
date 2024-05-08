package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"
)

type ReferensiArsitekturRepositoryImpl struct {
}

func NewReferensiArsitekturRepository() ReferensiArsitekturRepository {
	return &ReferensiArsitekturRepositoryImpl{}
}

func (repository *ReferensiArsitekturRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur) domain.ReferensiArsitektur {
	currentTime := time.Now()
	referensiarsitektur.Created_at = currentTime
	referensiarsitektur.Updated_at = currentTime

	script := "insert into referensi_arsitekturs(kode_referensi,nama_referensi,level_referensi,jenis_referensi,created_at,updated_at) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, referensiarsitektur.Kode_referensi, referensiarsitektur.Nama_referensi, referensiarsitektur.Level_referensi, referensiarsitektur.Jenis_referensi, referensiarsitektur.Created_at, referensiarsitektur.Updated_at)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	referensiarsitektur.IdReferensi = int(id)
	return referensiarsitektur
}

func (repository *ReferensiArsitekturRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur) domain.ReferensiArsitektur {
	currentTime := time.Now()
	referensiarsitektur.Updated_at = currentTime

	script := "update referensi_arsitekturs SET kode_referensi = ?, nama_referensi = ?, level_referensi = ?, jenis_referensi = ?, updated_at = ? where id_referensi = ?"
	_, err := tx.ExecContext(ctx, script, referensiarsitektur.Kode_referensi, referensiarsitektur.Nama_referensi, referensiarsitektur.Level_referensi, referensiarsitektur.Jenis_referensi, referensiarsitektur.Updated_at, referensiarsitektur.IdReferensi)
	helper.PanicIfError(err)

	return referensiarsitektur
}

func (repository *ReferensiArsitekturRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, referensiarsitektur domain.ReferensiArsitektur) {
	script := "delete from referensi_arsitekturs where id_referensi =?"
	_, err := tx.ExecContext(ctx, script, referensiarsitektur.IdReferensi)
	helper.PanicIfError(err)

}

func (repository *ReferensiArsitekturRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ReferensiArsitektur {
	script := "select id_referensi,kode_referensi,nama_referensi,level_referensi,jenis_referensi,created_at,updated_at from referensi_arsitekturs"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var reference []domain.ReferensiArsitektur
	for rows.Next() {
		referensi := domain.ReferensiArsitektur{}
		err := rows.Scan(&referensi.IdReferensi, &referensi.Kode_referensi, &referensi.Nama_referensi, &referensi.Level_referensi, &referensi.Jenis_referensi, &referensi.Created_at, &referensi.Updated_at)
		helper.PanicIfError(err)
		reference = append(reference, referensi)
	}
	return reference
}

func (repository *ReferensiArsitekturRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, referensiarsitekturId int) (domain.ReferensiArsitektur, error) {
	script := "select id_referensi, kode_referensi, nama_referensi, level_referensi, jenis_referensi, created_at, updated_at from referensi_arsitekturs where id_referensi = ?"
	rows, err := tx.QueryContext(ctx, script, referensiarsitekturId)
	helper.PanicIfError(err)
	defer rows.Close()

	referensiarsitektur := domain.ReferensiArsitektur{}
	if rows.Next() {
		err := rows.Scan(&referensiarsitektur.IdReferensi, &referensiarsitektur.Kode_referensi, &referensiarsitektur.Nama_referensi, &referensiarsitektur.Level_referensi, &referensiarsitektur.Jenis_referensi, &referensiarsitektur.Created_at, &referensiarsitektur.Updated_at)
		helper.PanicIfError(err)
		return referensiarsitektur, nil
	} else {
		return referensiarsitektur, errors.New("referensiarsitektur is not found")
	}
}

func (repository *ReferensiArsitekturRepositoryImpl) FindByKodeRef(ctx context.Context, tx *sql.Tx, kodeReferensi string) *helper.TreeNode {
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, "SELECT kode_referensi FROM referensi_arsitekturs WHERE kode_referensi LIKE ?", kodeReferensi+"%")
	} else {
		rows, err = tx.QueryContext(ctx, "SELECT kode_referensi FROM referensi_arsitekturs WHERE kode_referensi LIKE ?", kodeReferensi+"%")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	root := helper.NewTreeNode(kodeReferensi)

	for rows.Next() {
		var kodeReferensi string
		if err := rows.Scan(&kodeReferensi); err != nil {
			log.Fatal(err)
		}

		parts := strings.Split(kodeReferensi, ".")

		currNode := root

		for _, part := range parts {
			var foundChild *helper.TreeNode
			for _, child := range currNode.Children {
				if child.KodeReferensi == part {
					foundChild = child
					break
				}
			}

			if foundChild == nil {
				newChild := helper.NewTreeNode(part)
				currNode.AddChild(newChild)
				currNode = newChild
			} else {
				currNode = foundChild
			}
		}
	}

	return root
}
