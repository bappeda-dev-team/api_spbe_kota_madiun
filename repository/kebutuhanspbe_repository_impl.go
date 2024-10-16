package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type KebutuhanSPBERepositoryImpl struct {
}

func NewKebutuhanSPBERepositoryImpl() *KebutuhanSPBERepositoryImpl {
	return &KebutuhanSPBERepositoryImpl{}
}

func (repository *KebutuhanSPBERepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error) {
	script := "INSERT INTO kebutuhan_spbe(keterangan ,kode_opd, tahun, nama_domain, id_prosesbisnis, indikator_pj, pj) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, kebutuhanSPBE.Keterangan, kebutuhanSPBE.KodeOpd, kebutuhanSPBE.Tahun, kebutuhanSPBE.NamaDomain, kebutuhanSPBE.IdProsesbisnis, kebutuhanSPBE.IndikatorPj, kebutuhanSPBE.PenanggungJawab)
	if err != nil {
		return kebutuhanSPBE, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return kebutuhanSPBE, err
	}
	kebutuhanSPBE.ID = int(id)

	for i := range kebutuhanSPBE.JenisKebutuhan {
		kebutuhanSPBE.JenisKebutuhan[i].KebutuhanId = kebutuhanSPBE.ID
		scriptJK := "INSERT INTO jenis_kebutuhan(kebutuhan_id, kebutuhan) VALUES (?, ?)"
		resultJK, err := tx.ExecContext(ctx, scriptJK, kebutuhanSPBE.JenisKebutuhan[i].KebutuhanId, kebutuhanSPBE.JenisKebutuhan[i].Kebutuhan)
		if err != nil {
			return kebutuhanSPBE, err
		}

		jkId, err := resultJK.LastInsertId()
		if err != nil {
			return kebutuhanSPBE, err
		}
		kebutuhanSPBE.JenisKebutuhan[i].Id = int(jkId)

		for j := range kebutuhanSPBE.JenisKebutuhan[i].KondisiAwal {
			kebutuhanSPBE.JenisKebutuhan[i].KondisiAwal[j].JenisKebutuhanId = kebutuhanSPBE.JenisKebutuhan[i].Id
			scriptKA := "INSERT INTO kondisi_awal(jenis_kebutuhan_id, keterangan, tahun ) VALUES (?, ?, ?)"
			resultKA, err := tx.ExecContext(ctx, scriptKA, kebutuhanSPBE.JenisKebutuhan[i].KondisiAwal[j].JenisKebutuhanId, kebutuhanSPBE.JenisKebutuhan[i].KondisiAwal[j].Keterangan, kebutuhanSPBE.JenisKebutuhan[i].KondisiAwal[j].Tahun)
			if err != nil {
				return kebutuhanSPBE, err
			}

			kaId, err := resultKA.LastInsertId()
			if err != nil {
				return kebutuhanSPBE, err
			}
			kebutuhanSPBE.JenisKebutuhan[i].KondisiAwal[j].Id = int(kaId)
		}
	}

	return kebutuhanSPBE, nil
}

func (repository *KebutuhanSPBERepositoryImpl) FindByKodeOpdAndTahun(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int, prosesbisnis int) ([]domain.KebutuhanSPBE, error) {
	log.Println("Menjalankan FindByKodeOpdAndTahun dengan kodeOpd:", kodeOpd, "dan tahun:", tahun, "dan prosesbisnisId:", prosesbisnis)

	script := "SELECT id, keterangan, kode_opd, tahun, nama_domain, id_prosesbisnis, indikator_pj, pj FROM kebutuhan_spbe WHERE 1=1"
	var args []interface{}

	if kodeOpd != "" {
		script += " AND kode_opd = ?"
		args = append(args, kodeOpd)
	}
	if tahun != 0 {
		script += " AND tahun = ?"
		args = append(args, tahun)
	}
	if prosesbisnis != 0 {
		script += " AND id_prosesbisnis =?"
		args = append(args, prosesbisnis)
	}

	script += " ORDER BY id DESC"

	log.Println("Menjalankan query:", script, "dengan args:", args)

	rows, err := tx.QueryContext(ctx, script, args...)
	if err != nil {
		log.Println("Error saat menjalankan query:", err)
		return nil, err
	}
	defer rows.Close()

	var kebutuhanSPBEs []domain.KebutuhanSPBE
	for rows.Next() {
		kebutuhanSPBE := domain.KebutuhanSPBE{}
		err := rows.Scan(&kebutuhanSPBE.ID, &kebutuhanSPBE.Keterangan, &kebutuhanSPBE.KodeOpd, &kebutuhanSPBE.Tahun, &kebutuhanSPBE.NamaDomain, &kebutuhanSPBE.IdProsesbisnis, &kebutuhanSPBE.IndikatorPj, &kebutuhanSPBE.PenanggungJawab)
		if err != nil {
			log.Println("Error saat memindai baris:", err)
			return nil, err
		}
		kebutuhanSPBEs = append(kebutuhanSPBEs, kebutuhanSPBE)
	}
	return kebutuhanSPBEs, nil
}

func (repository *KebutuhanSPBERepositoryImpl) FindJenisKebutuhanByKebutuhanId(ctx context.Context, tx *sql.Tx, kebutuhanId int) ([]domain.JenisKebutuhan, error) {
	script := "SELECT id, kebutuhan_id, kebutuhan FROM jenis_kebutuhan WHERE kebutuhan_id = ?"
	rows, err := tx.QueryContext(ctx, script, kebutuhanId)
	if err != nil {
		log.Println("Error saat menjalankan query jenis_kebutuhan:", err)
		return nil, err
	}
	defer rows.Close()

	var jenisKebutuhans []domain.JenisKebutuhan
	for rows.Next() {
		jk := domain.JenisKebutuhan{}
		err := rows.Scan(&jk.Id, &jk.KebutuhanId, &jk.Kebutuhan)
		if err != nil {
			log.Println("Error saat memindai baris jenis_kebutuhan:", err)
			return nil, err
		}
		jenisKebutuhans = append(jenisKebutuhans, jk)
	}
	return jenisKebutuhans, nil
}

func (repository *KebutuhanSPBERepositoryImpl) FindKondisiAwalByJenisKebutuhanId(ctx context.Context, tx *sql.Tx, jenisKebutuhanId int) ([]domain.KondisiAwal, error) {
	script := "SELECT id, jenis_kebutuhan_id, keterangan, tahun FROM kondisi_awal WHERE jenis_kebutuhan_id = ?"
	rows, err := tx.QueryContext(ctx, script, jenisKebutuhanId)
	if err != nil {
		log.Println("Error saat menjalankan query kondisi_awal:", err)
		return nil, err
	}
	defer rows.Close()

	var kondisiAwals []domain.KondisiAwal
	for rows.Next() {
		ka := domain.KondisiAwal{}
		err := rows.Scan(&ka.Id, &ka.JenisKebutuhanId, &ka.Keterangan, &ka.Tahun)
		if err != nil {
			log.Println("Error saat memindai baris kondisi_awal:", err)
			return nil, err
		}
		kondisiAwals = append(kondisiAwals, ka)
	}
	return kondisiAwals, nil
}

func (repository *KebutuhanSPBERepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kebutuhanSPBEId int) (domain.KebutuhanSPBE, error) {
	log.Println("Menjalankan FindById dengan kebutuhanSPBEId:", kebutuhanSPBEId)

	script := "SELECT id, keterangan, kode_opd, tahun, nama_domain, id_prosesbisnis, indikator_pj, pj FROM kebutuhan_spbe WHERE id = ?"
	args := []interface{}{kebutuhanSPBEId}

	log.Println("Menjalankan query:", script, "dengan args:", args)

	row := tx.QueryRowContext(ctx, script, args...)
	kebutuhanSPBE := domain.KebutuhanSPBE{}
	err := row.Scan(&kebutuhanSPBE.ID, &kebutuhanSPBE.Keterangan, &kebutuhanSPBE.KodeOpd, &kebutuhanSPBE.Tahun, &kebutuhanSPBE.NamaDomain, &kebutuhanSPBE.IdProsesbisnis, &kebutuhanSPBE.IndikatorPj, &kebutuhanSPBE.PenanggungJawab)
	if err != nil {
		log.Println("Error saat memindai baris:", err)
		return kebutuhanSPBE, err
	}

	return kebutuhanSPBE, nil
}

func (repository *KebutuhanSPBERepositoryImpl) Update(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error) {
	script := "UPDATE kebutuhan_spbe SET kode_opd = ?, tahun = ?, nama_domain = ?, id_prosesbisnis = ?, indikator_pj = ?, pj = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, kebutuhanSPBE.KodeOpd, kebutuhanSPBE.Tahun, kebutuhanSPBE.NamaDomain, kebutuhanSPBE.IdProsesbisnis, kebutuhanSPBE.IndikatorPj, kebutuhanSPBE.PenanggungJawab, kebutuhanSPBE.ID)
	if err != nil {
		log.Println("Error saat mengupdate kebutuhan_spbe:", err)
		return kebutuhanSPBE, err
	}

	// Hapus jenis kebutuhan dan kondisi awal yang ada
	_, err = tx.ExecContext(ctx, "DELETE FROM kondisi_awal WHERE jenis_kebutuhan_id IN (SELECT id FROM jenis_kebutuhan WHERE kebutuhan_id = ?)", kebutuhanSPBE.ID)
	if err != nil {
		log.Println("Error saat menghapus kondisi_awal:", err)
		return kebutuhanSPBE, err
	}
	_, err = tx.ExecContext(ctx, "DELETE FROM jenis_kebutuhan WHERE kebutuhan_id = ?", kebutuhanSPBE.ID)
	if err != nil {
		log.Println("Error saat menghapus jenis_kebutuhan:", err)
		return kebutuhanSPBE, err
	}

	// Masukkan jenis kebutuhan dan kondisi awal yang baru
	for _, jk := range kebutuhanSPBE.JenisKebutuhan {
		scriptJK := "INSERT INTO jenis_kebutuhan(kebutuhan_id, kebutuhan) VALUES (?, ?)"
		resultJK, err := tx.ExecContext(ctx, scriptJK, kebutuhanSPBE.ID, jk.Kebutuhan)
		if err != nil {
			log.Println("Error saat memasukkan jenis_kebutuhan:", err)
			return kebutuhanSPBE, err
		}

		jkId, err := resultJK.LastInsertId()
		if err != nil {
			log.Println("Error saat mendapatkan LastInsertId untuk jenis_kebutuhan:", err)
			return kebutuhanSPBE, err
		}

		for _, ka := range jk.KondisiAwal {
			scriptKA := "INSERT INTO kondisi_awal(jenis_kebutuhan_id, keterangan, tahun) VALUES (?, ?, ?)"
			_, err := tx.ExecContext(ctx, scriptKA, jkId, ka.Keterangan, ka.Tahun)
			if err != nil {
				log.Println("Error saat memasukkan kondisi_awal:", err)
				return kebutuhanSPBE, err
			}
		}
	}

	return kebutuhanSPBE, nil
}

func (repository *KebutuhanSPBERepositoryImpl) UpdateKeterangan(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error) {
	script := "UPDATE kebutuhan_spbe SET keterangan = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, kebutuhanSPBE.Keterangan, kebutuhanSPBE.ID)
	if err != nil {
		log.Println("Error saat mengupdate kebutuhan_spbe:", err)
		return kebutuhanSPBE, err
	}
	return kebutuhanSPBE, nil
}

func (repository *KebutuhanSPBERepositoryImpl) UpdatePenanggungJawab(ctx context.Context, tx *sql.Tx, kebutuhanSPBE domain.KebutuhanSPBE) (domain.KebutuhanSPBE, error) {
	script := "UPDATE kebutuhan_spbe SET indikator_pj = ?, pj = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, kebutuhanSPBE.IndikatorPj, kebutuhanSPBE.PenanggungJawab, kebutuhanSPBE.ID)
	if err != nil {
		log.Println("Error saat mengupdate kebutuhan_spbe:", err)
		return kebutuhanSPBE, err
	}
	return kebutuhanSPBE, nil
}

func (repository *KebutuhanSPBERepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, kebutuhanSPBEId int) error {
	// Hapus kondisi_awal yang terhubung
	_, err := tx.ExecContext(ctx, "DELETE FROM kondisi_awal WHERE jenis_kebutuhan_id IN (SELECT id FROM jenis_kebutuhan WHERE kebutuhan_id = ?)", kebutuhanSPBEId)
	if err != nil {
		return err
	}

	// Hapus jenis_kebutuhan yang terhubung
	_, err = tx.ExecContext(ctx, "DELETE FROM jenis_kebutuhan WHERE kebutuhan_id = ?", kebutuhanSPBEId)
	if err != nil {
		return err
	}

	// Hapus kebutuhan_spbe
	script := "DELETE FROM kebutuhan_spbe WHERE id = ?"
	_, err = tx.ExecContext(ctx, script, kebutuhanSPBEId)
	return err
}

func (repository *KebutuhanSPBERepositoryImpl) FindPenanggungJawab(ctx context.Context, tx *sql.Tx, pj string, tahun int) ([]domain.KebutuhanSPBE, error) {
	log.Println("Menjalankan Find Penanggung jawab dengan pj:", pj)

	script := "SELECT id, keterangan, kode_opd, tahun, nama_domain, id_prosesbisnis, indikator_pj, pj FROM kebutuhan_spbe WHERE 1=1"
	var args []interface{}

	if pj != "" {
		script += " AND pj = ?"
		args = append(args, pj)
	}

	if tahun != 0 {
		script += " AND tahun = ?"
		args = append(args, tahun)
	}

	script += " ORDER BY id DESC"

	log.Println("Menjalankan query:", script, "dengan args:", args)

	rows, err := tx.QueryContext(ctx, script, args...)
	if err != nil {
		log.Println("Error saat menjalankan query:", err)
		return nil, err
	}
	defer rows.Close()

	var kebutuhanSPBEs []domain.KebutuhanSPBE
	for rows.Next() {
		kebutuhanSPBE := domain.KebutuhanSPBE{}
		err := rows.Scan(&kebutuhanSPBE.ID, &kebutuhanSPBE.Keterangan, &kebutuhanSPBE.KodeOpd, &kebutuhanSPBE.Tahun, &kebutuhanSPBE.NamaDomain, &kebutuhanSPBE.IdProsesbisnis, &kebutuhanSPBE.IndikatorPj, &kebutuhanSPBE.PenanggungJawab)
		if err != nil {
			log.Println("Error saat memindai baris:", err)
			return nil, err
		}
		kebutuhanSPBEs = append(kebutuhanSPBEs, kebutuhanSPBE)
	}
	return kebutuhanSPBEs, nil
}

func (repository *KebutuhanSPBERepositoryImpl) FindIdPenanggungJawab(ctx context.Context, tx *sql.Tx, kebutuhanId int, pj string) (domain.KebutuhanSPBE, error) {
	log.Printf("Mencari kebutuhan SPBE dengan ID: %d dan PJ: %s", kebutuhanId, pj)

	script := "SELECT id, keterangan, kode_opd, tahun, nama_domain, id_prosesbisnis, indikator_pj, pj FROM kebutuhan_spbe WHERE 1=1"
	var args []interface{}

	if kebutuhanId != 0 {
		script += " AND id = ?"
		args = append(args, kebutuhanId)
	}

	if pj != "" {
		script += " AND pj = ?"
		args = append(args, pj)
	}

	log.Println("Menjalankan query:", script, "dengan args:", args)

	row := tx.QueryRowContext(ctx, script, args...)
	kebutuhanSPBE := domain.KebutuhanSPBE{}
	err := row.Scan(&kebutuhanSPBE.ID, &kebutuhanSPBE.Keterangan, &kebutuhanSPBE.KodeOpd, &kebutuhanSPBE.Tahun, &kebutuhanSPBE.NamaDomain, &kebutuhanSPBE.IdProsesbisnis, &kebutuhanSPBE.IndikatorPj, &kebutuhanSPBE.PenanggungJawab)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Hasil query: kebutuhan SPBE tidak ditemukan")
			return domain.KebutuhanSPBE{}, fmt.Errorf("kebutuhan SPBE tidak ditemukan")
		}
		log.Println("Error saat memindai baris:", err)
		return domain.KebutuhanSPBE{}, fmt.Errorf("error saat memindai baris: %v", err)
	}

	log.Printf("Hasil query: %+v", kebutuhanSPBE)
	return kebutuhanSPBE, nil
}
