CREATE TABLE urusan
(
    id INT NOT NULL AUTO_INCREMENT,
    kode_urusan VARCHAR(255) UNIQUE,
    urusan VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX(kode_urusan)
) ENGINE=InnoDB;
