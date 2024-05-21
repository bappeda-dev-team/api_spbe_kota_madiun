CREATE TABLE proses_bisnis
(
    id INT NOT NULL AUTO_INCREMENT,
    nama_proses_bisnis TEXT NOT NULL,
    sasaran_kota TEXT NOT NULL,
    kode_proses_bisinis TEXT NOT NULL,
    kode_opd TEXT NOT NULL,
    bidang_urusan TEXT NOT NULL,
    rad_level_1 TEXT NOT NULL,
    rad_level_2 TEXT NOT NULL,
    rad_level_3 TEXT NOT NULL,
    rad_level_4 TEXT NOT NULL,
    rad_level_5 TEXT NOT NULL,
    rad_level_6 TEXT NOT NULL,
    tahun INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB;
