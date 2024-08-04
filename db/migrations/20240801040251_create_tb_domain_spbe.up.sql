CREATE TABLE domain_spbe
(
    id INT NOT NULL AUTO_INCREMENT,
    nama_domain VARCHAR(255) UNIQUE,
    kode_domain VARCHAR(255),
    tahun INT,
    PRIMARY KEY (id),
    INDEX (nama_domain)
) ENGINE=InnoDB;