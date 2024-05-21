package repository

import (
	"api_spbe_kota_madiun/helper"
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ProsesBisnisRepositoryImpl struct {
}

func NewProsesBisnisRepository() ProsesBisnisRepository {
	return &ProsesBisnisRepositoryImpl{}
}

func (reporepository *ProsesBisnisRepositoryImpl)Save(ctx context.Context, tx *sql.Tx, prosesbisnis domain.ProsesBisnis)domain.ProsesBisnis{
	currentTime := time.Now()
	prosesbisnis.Created_at = currentTime
	prosesbisnis.Updated_at = currentTime

	script := "insert into proses_bisnis(nama_proses,sasaran_kota,kode_proses_bisnis,kode_opd,bidang_urusan,rad_level_1,rad_level_2,rad_level_3,rad_level_4,rad_level_5,rad_level_6,created_at,updated_at) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, prosesbisnis.Nama_proses_bisnis, prosesbisnis.Sasaran_kota, prosesbisnis.Kode_proses_bisnis, prosesbisnis.Kode_opd, prosesbisnis.Bidang_urusan,prosesbisnis.Rad_level_1,prosesbisnis.Rad_level_2,prosesbisnis.Rad_level_3,prosesbisnis.Rad_level_4,prosesbisnis.Rad_level_5,prosesbisnis.Rad_level_6,prosesbisnis.Created_at, prosesbisnis.Updated_at)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	prosesbisnis.Id = int(id)
	return prosesbisnis
}
func (reporepository *ProsesBisnisRepositoryImpl)Update(ctx context.Context, tx *sql.Tx, prosesbisnis domain.ProsesBisnis)domain.ProsesBisnis{
	currentTime := time.Now()
	prosesbisnis.Updated_at = currentTime

	script := "insert into proses_bisnis(nama_proses,sasaran_kota,kode_proses_bisnis,kode_opd,bidang_urusan,rad_level_1,rad_level_2,rad_level_3,rad_level_4,rad_level_5,rad_level_6,created_at,updated_at) values (?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, script, prosesbisnis.Nama_proses_bisnis, prosesbisnis.Sasaran_kota, prosesbisnis.Kode_proses_bisnis, prosesbisnis.Kode_opd, prosesbisnis.Bidang_urusan,prosesbisnis.Rad_level_1,prosesbisnis.Rad_level_2,prosesbisnis.Rad_level_3,prosesbisnis.Rad_level_4,prosesbisnis.Rad_level_5,prosesbisnis.Rad_level_6,prosesbisnis.Created_at, prosesbisnis.Updated_at)
	helper.PanicIfError(err)

	return prosesbisnis
}

func (reporepository *ProsesBisnisRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, prosesbisnis domain.ProsesBisnis){
	script := "delete from proses_bisnis where id =?"
	_, err := tx.ExecContext(ctx, script, prosesbisnis.Id)
	helper.PanicIfError(err)
}

func (reporepository *ProsesBisnisRepositoryImpl)FindById(ctx context.Context, tx *sql.Tx, Id int)(domain.ProsesBisnis,error){
	script := "select id,nama_proses,sasaran_kota,kode_proses_bisnis,kode_opd,bidang_urusan,rad_level_1,rad_level_2,rad_level_3,rad_level_4,rad_level_5,rad_level_6,created_at,updated_at where id = ?"
	rows, err := tx.QueryContext(ctx, script, Id)
	helper.PanicIfError(err)
	defer rows.Close()

	prosesbisnis := domain.ProsesBisnis{}
	if rows.Next() {
		err := rows.Scan(&prosesbisnis.Id,&prosesbisnis.Nama_proses_bisnis, &prosesbisnis.Sasaran_kota, &prosesbisnis.Kode_proses_bisnis, &prosesbisnis.Kode_opd, &prosesbisnis.Bidang_urusan, &prosesbisnis.Rad_level_1, &prosesbisnis.Rad_level_2, &prosesbisnis.Rad_level_3, &prosesbisnis.Rad_level_4,&prosesbisnis.Rad_level_5,&prosesbisnis.Rad_level_6,&prosesbisnis.Created_at, &prosesbisnis.Updated_at)
		helper.PanicIfError(err)
		return prosesbisnis, nil
	} else {
		return prosesbisnis, errors.New("proses bisnis is not found")
	}
}

func (reporepository *ProsesBisnisRepositoryImpl)FindByKodeOdp(ctx context.Context, tx *sql.Tx, kode_opd string)(domain.ProsesBisnis, error){
	script := "select id,nama_proses,sasaran_kota,kode_proses_bisnis,kode_opd,bidang_urusan,rad_level_1,rad_level_2,rad_level_3,rad_level_4,rad_level_5,rad_level_6,created_at,updated_at where kode_opd = ?"
	rows, err := tx.QueryContext(ctx, script, kode_opd)
	helper.PanicIfError(err)
	defer rows.Close()

	prosesbisnis := domain.ProsesBisnis{}
	if rows.Next() {
		err := rows.Scan(&prosesbisnis.Id,&prosesbisnis.Nama_proses_bisnis, &prosesbisnis.Sasaran_kota, &prosesbisnis.Kode_proses_bisnis, &prosesbisnis.Kode_opd, &prosesbisnis.Bidang_urusan, &prosesbisnis.Rad_level_1, &prosesbisnis.Rad_level_2, &prosesbisnis.Rad_level_3, &prosesbisnis.Rad_level_4,&prosesbisnis.Rad_level_5,&prosesbisnis.Rad_level_6,&prosesbisnis.Created_at, &prosesbisnis.Updated_at)
		helper.PanicIfError(err)
		return prosesbisnis, nil
	} else {
		return prosesbisnis, errors.New("Kode Opd is not found")
	}
}

func (reporepository *ProsesBisnisRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx)[]domain.ProsesBisnis{
	script := "select id,nama_proses,sasaran_kota,kode_proses_bisnis,kode_opd,bidang_urusan,rad_level_1,rad_level_2,rad_level_3,rad_level_4,rad_level_5,rad_level_6,created_at,updated_at from proses_bisnis"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()
	
	var proses []domain.ProsesBisnis
	for rows.Next() {
		prosesbisnis := domain.ProsesBisnis{}
		err := rows.Scan(&prosesbisnis.Id,&prosesbisnis.Nama_proses_bisnis, &prosesbisnis.Sasaran_kota, &prosesbisnis.Kode_proses_bisnis, &prosesbisnis.Kode_opd, &prosesbisnis.Bidang_urusan, &prosesbisnis.Rad_level_1, &prosesbisnis.Rad_level_2, &prosesbisnis.Rad_level_3, &prosesbisnis.Rad_level_4,&prosesbisnis.Rad_level_5,&prosesbisnis.Rad_level_6,&prosesbisnis.Created_at, &prosesbisnis.Updated_at)
		helper.PanicIfError(err)
		proses = append(proses, prosesbisnis)
	}
	return proses


}