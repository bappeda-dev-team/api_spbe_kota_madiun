CREATE TABLE pohon_kinerja
(
    id INT NOT NULL AUTO_INCREMENT,
    nama_pohon VARCHAR(255) UNIQUE,
    jenis_pohon VARCHAR(255),
    level_pohon INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    tahun INT,
    kode_opd VARCHAR(255),
    PRIMARY KEY (id),
    CONSTRAINT fk_pohon_kinerja_kode_opd FOREIGN KEY (kode_opd) REFERENCES opd (kode_opd)
) ENGINE=InnoDB;
