package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type PetarencanaRepositoryImpl struct {
}

func NewPetarencanaRepositoryImpl() *PetarencanaRepositoryImpl {
	return &PetarencanaRepositoryImpl{}
}

func (repository *PetarencanaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun int) ([]domain.Petarencana, error) {
	query := `
	SELECT
	pb.id,
	pb.kode_opd,
	pb.tahun,
	pb.nama_proses_bisnis,
	pb.kode_proses_bisnis,
	l.nama_layanan,
	d.nama_data,
	a.nama_aplikasi,
	ks.id AS id_keterangan
	FROM
	proses_bisnis pb
	LEFT JOIN
	layanan_spbe l ON (l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id AND l.operational_id = pb.operational_id AND l.kode_opd = pb.kode_opd)
	OR (l.strategic_id = pb.strategic_id AND l.tactical_id = pb.tactical_id AND l.kode_opd = pb.kode_opd)
	OR (l.strategic_id = pb.strategic_id AND l.kode_opd = pb.kode_opd)
	LEFT JOIN
	data_dan_informasi d ON (d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id AND d.operational_id = pb.operational_id AND d.kode_opd = pb.kode_opd)
	OR (d.strategic_id = pb.strategic_id AND d.tactical_id = pb.tactical_id AND d.kode_opd = pb.kode_opd)
	OR (d.strategic_id = pb.strategic_id AND d.kode_opd = pb.kode_opd)
	LEFT JOIN
	aplikasi a ON (a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id AND a.operational_id = pb.operational_id AND a.kode_opd = pb.kode_opd)
	OR (a.strategic_id = pb.strategic_id AND a.tactical_id = pb.tactical_id AND a.kode_opd = pb.kode_opd)
	OR (a.strategic_id = pb.strategic_id AND a.kode_opd = pb.kode_opd)
	LEFT JOIN
	kebutuhan_spbe ks ON ks.id_prosesbisnis = pb.id
	WHERE 1=1
	AND l.nama_layanan IS NOT NULL
	AND d.nama_data IS NOT NULL
	AND a.nama_aplikasi IS NOT NULL

	`

	var args []interface{}
	if kodeOpd != "" {
		query += " AND pb.kode_opd = ?"
		args = append(args, kodeOpd)
	}
	if tahun != 0 {
		query += " AND pb.tahun = ?"
		args = append(args, tahun)
	}

	query += " ORDER BY pb.kode_opd DESC, pb.id DESC;"

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var petarencanaList []domain.Petarencana
	for rows.Next() {
		var id int
		var kodeOpd string
		var tahun int
		var namaProsesBisnis, kodeProsesBisnis string
		var namaLayanan, namaData, namaAplikasi sql.NullString
		var idKeterangan sql.NullInt32

		if err := rows.Scan(
			&id,
			&kodeOpd,
			&tahun,
			&namaProsesBisnis,
			&kodeProsesBisnis,
			&namaLayanan,
			&namaData,
			&namaAplikasi,
			&idKeterangan,
		); err != nil {
			return nil, err
		}

		petarencana := domain.Petarencana{
			ID:               id,
			KodeOpd:          kodeOpd,
			Tahun:            tahun,
			NamaProsesBisnis: namaProsesBisnis,
			KodeProsesBisnis: kodeProsesBisnis,
			Layanan: []domain.RencanaLayanan{
				{
					NamaLayanan: namaLayanan,
				},
			},
			DataDanInformasi: []domain.RencanaDataDanInformasi{
				{
					NamaData: namaData,
				},
			},
			Aplikasi: []domain.RencanaAplikasi{
				{
					NamaAplikasi: namaAplikasi,
				},
			},
			Keterangan: []domain.RencanaKeterangan{
				{
					IdKeterangan: idKeterangan,
				},
			},
		}

		petarencanaList = append(petarencanaList, petarencana)
	}

	return petarencanaList, nil
}
