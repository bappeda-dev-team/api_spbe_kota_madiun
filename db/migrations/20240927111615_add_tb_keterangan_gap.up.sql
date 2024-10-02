CREATE TABLE keterangan_gap
(
    id INT NOT NULL AUTO_INCREMENT,
    keterangan_gap VARCHAR(255),
    kode_opd VARCHAR(255),
    id_prosesbisnis INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_keterangan_gap_kode_opd FOREIGN KEY (kode_opd) REFERENCES opd (kode_opd),
    CONSTRAINT fk_keterangan_gap_id_prosesbisnis FOREIGN KEY (id_prosesbisnis) REFERENCES proses_bisnis (id)
) ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;