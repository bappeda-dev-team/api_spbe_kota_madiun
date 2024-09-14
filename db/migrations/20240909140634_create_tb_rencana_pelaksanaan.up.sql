CREATE TABLE rencana_pelaksanaan (
    id INT NOT NULL AUTO_INCREMENT,
    id_kebutuhan INT, 
    kode_opd VARCHAR(255),
    id_sasarankinerja INT,
    indikator_perangkatdaerah ENUM('internal', 'eksternal'),
    perangkat_daerah VARCHAR(255),
    PRIMARY KEY (id),
    CONSTRAINT fk_kebutuhan_spbe FOREIGN KEY (id_kebutuhan) REFERENCES kebutuhan_spbe (id),
    CONSTRAINT fk_sasarankinerja_pegawai FOREIGN KEY (id_sasarankinerja) REFERENCES sasaran_kinerja_pegawai (id),
    CONSTRAINT fk_perangkatdaerah FOREIGN KEY (perangkat_daerah) REFERENCES opd (kode_opd),
    CONSTRAINT fk_kode_opd_rencana FOREIGN KEY (kode_opd) REFERENCES opd (kode_opd)
)ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;