CREATE TABLE kebutuhan_spbe
(
    id INT NOT NULL AUTO_INCREMENT,
    nama_domain VARCHAR(255),
    kode_opd VARCHAR(255),
    tahun INT NOT NULL,
    id_prosesbisnis INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_domain FOREIGN KEY (nama_domain) REFERENCES domain_spbe (nama_domain),
    CONSTRAINT fk_kebutuhan_spbe_kode_opd FOREIGN KEY (kode_opd) REFERENCES opd (kode_opd),
    CONSTRAINT fk_kebutuhan_spbe_id_prosesbisnis FOREIGN KEY (id_prosesbisnis) REFERENCES proses_bisnis (id)
) ENGINE=InnoDB
    DEFAULT CHARSET=latin1
    COLLATE latin1_swedish_ci;
