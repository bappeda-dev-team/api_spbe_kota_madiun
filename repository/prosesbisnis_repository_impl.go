package repository

import (
	"api_spbe_kota_madiun/model/domain"
	"context"
	"database/sql"
)

type ProsesBisnisRepositoryImpl struct {
}

func NewProsesBisnisRepository() ProsesBisnisRepository {
	return &ProsesBisnisRepositoryImpl{}
}


func (repository *ProsesBisnisRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx)([]domain.ProsesBisnis, error){
	script := "SELECT id, nama_proses_bisnis, sasaran_kota, kode_proses_bisnis, kode_opd, bidang_urusan, rab_level_1_id, rab_level_2_id, rab_level_3_id, rab_level_4_id, rab_level_5_id, rab_level_6_id, tahun, created_at, updated_at FROM proses_bisnis"
	rows, err := tx.Query(script)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var prosesBisnisList []domain.ProsesBisnis
    for rows.Next() {
        var prosesBisnis domain.ProsesBisnis
        err := rows.Scan(&prosesBisnis.ID, &prosesBisnis.NamaProsesBisnis, &prosesBisnis.SasaranKota, &prosesBisnis.KodeProsesBisnis, &prosesBisnis.KodeOPD, &prosesBisnis.BidangUrusan, &prosesBisnis.RabLevel1ID, &prosesBisnis.RabLevel2ID, &prosesBisnis.RabLevel3ID, &prosesBisnis.RabLevel4ID, &prosesBisnis.RabLevel5ID, &prosesBisnis.RabLevel6ID, &prosesBisnis.Tahun, &prosesBisnis.CreatedAt, &prosesBisnis.UpdatedAt)
        if err != nil {
            return nil, err
        }
        prosesBisnisList = append(prosesBisnisList, prosesBisnis)
    }
    return prosesBisnisList, nil
}