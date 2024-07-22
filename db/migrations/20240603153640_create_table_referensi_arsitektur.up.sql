CREATE TABLE referensi_arsitekturs
(
    id INT NOT NULL AUTO_INCREMENT,
    kode_referensi VARCHAR(255) NOT NULL,
    nama_referensi VARCHAR(255) NOT NULL,
    level_referensi INT NOT NULL,
    jenis_referensi VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    tahun INT(24) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB;
