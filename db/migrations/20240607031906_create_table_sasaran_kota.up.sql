CREATE TABLE sasaran_kota
(
    id INT NOT NULL AUTO_INCREMENT,
    sasaran VARCHAR(255) UNIQUE,
    strategi_kota VARCHAR(255),
    tujuan_kota VARCHAR(255),
    tahun INT(24),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (sasaran)
) ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;