CREATE TABLE proses_bisnis
(
    id INT NOT NULL AUTO_INCREMENT,
    nama_proses_bisnis TEXT NOT NULL,
    sasaran_kota VARCHAR(255) NOT NULL,
    kode_proses_bisnis VARCHAR(255) NOT NULL,
    kode_opd VARCHAR(255) NOT NULL,
    bidang_urusan VARCHAR(255) NOT NULL,
    rab_level_1_id INT,
    rab_level_2_id INT,
    rab_level_3_id INT,
    rab_level_4_id INT,
    rab_level_5_id INT,
    rab_level_6_id INT,
    tahun INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_rab_level_1 FOREIGN KEY (rab_level_1_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_rab_level_2 FOREIGN KEY (rab_level_2_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_rab_level_3 FOREIGN KEY (rab_level_3_id) REFERENCES referensi_arsitekturs (id)
) ENGINE=InnoDB;
