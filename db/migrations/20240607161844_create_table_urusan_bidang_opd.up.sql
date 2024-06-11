CREATE TABLE urusan_bidang_opd
(
    id INT NOT NULL AUTO_INCREMENT,
    kode_opd VARCHAR(255),
    kode_urusan VARCHAR(255),
    bidang_urusan VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_kode_urusan FOREIGN KEY(kode_urusan) REFERENCES urusan(kode_urusan),
    CONSTRAINT fk_bidang_urusan FOREIGN KEY(bidang_urusan) REFERENCES bidang_urusan(kode_bidang_urusan),
    CONSTRAINT fk_kode_opd FOREIGN KEY (kode_opd) REFERENCES opd(kode_opd)
) ENGINE=InnoDB;
