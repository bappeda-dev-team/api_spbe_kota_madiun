CREATE TABLE layanan_spbe
(
    id INT NOT NULL AUTO_INCREMENT,
    nama_layanan TEXT,
    kode_layanan VARCHAR(255),
    tujuan_layanan_id INT,
    fungsi_layanan VARCHAR(255),
    tahun INT NOT NULL,
    kode_opd VARCHAR(255),
    kementrian_terkait VARCHAR(255),
    metode_layanan VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    ral_level_1_id INT,
    ral_level_2_id INT,
    ral_level_3_id INT,
    ral_level_4_id INT,
    ral_level_5_id INT,
    PRIMARY KEY (id),
    CONSTRAINT fk_ral_level_1 FOREIGN KEY (ral_level_1_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_ral_level_2 FOREIGN KEY (ral_level_2_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_ral_level_3 FOREIGN KEY (ral_level_3_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_ral_level_4 FOREIGN KEY (ral_level_4_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_ral_level_5 FOREIGN KEY (ral_level_5_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_tujuan_layanan FOREIGN KEY (tujuan_layanan_id) REFERENCES pohon_kinerja (id),
    CONSTRAINT fk_layanan_spbe_kode_opd FOREIGN KEY (kode_opd) REFERENCES opd (kode_opd)
) ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;