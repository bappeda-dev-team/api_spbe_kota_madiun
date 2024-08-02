CREATE TABLE kondisi_awal
(
    id INT NOT NULL AUTO_INCREMENT,
    tahun INT,
    keterangan TEXT,
    jenis_kebutuhan_id INT,
    PRIMARY KEY (id),
    CONSTRAINT fk_jenis_kebutuhan_id FOREIGN KEY (jenis_kebutuhan_id) REFERENCES jenis_kebutuhan (id)
) ENGINE=InnoDB;