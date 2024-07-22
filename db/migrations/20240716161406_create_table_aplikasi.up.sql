CREATE TABLE aplikasi(
    id INT NOT NULL AUTO_INCREMENT,
    nama_aplikasi TEXT,
    fungsi_aplikasi VARCHAR(255),
    jenis_aplikasi VARCHAR(255),
    produsen_aplikasi VARCHAR(255),
    pj_aplikasi VARCHAR(255),
    informasi_terkait_input VARCHAR(255),
    informasi_terkait_output VARCHAR(255),
    interoprabilitas VARCHAR(255),
    kode_opd varchar(255),
    tahun INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    raa_level_1_id INT,
    raa_level_2_id INT,
    raa_level_3_id INT,
    raa_level_4_id INT,
    strategic_id INT,
    tactical_id INT,
    operational_id INT,
    PRIMARY KEY (id),
    CONSTRAINT fk_raa_level_1 FOREIGN KEY (raa_level_1_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_raa_level_2 FOREIGN KEY (raa_level_2_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_raa_level_3 FOREIGN KEY (raa_level_3_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_raa_level_4 FOREIGN KEY (raa_level_4_id) REFERENCES referensi_arsitekturs (id),
    CONSTRAINT fk_strategic_aplikasi FOREIGN KEY (strategic_id) REFERENCES pohon_kinerja (id),
    CONSTRAINT fk_tactical_aplikasi FOREIGN KEY (tactical_id) REFERENCES pohon_kinerja (id),
    CONSTRAINT fk_operational_aplikasi FOREIGN KEY (operational_id) REFERENCES pohon_kinerja (id),
    CONSTRAINT fk_kode_opd_aplikasi FOREIGN KEY (kode_opd) REFERENCES opd (kode_opd)
)ENGINE=InnoDB;