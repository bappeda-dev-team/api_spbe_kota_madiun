CREATE TABLE jenis_kebutuhan
(
    id INT NOT NULL AUTO_INCREMENT,
    kebutuhan_id INT,
    kebutuhan TEXT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_kebutuhan_id FOREIGN KEY (kebutuhan_id) REFERENCES kebutuhan_spbe (id)
) ENGINE=InnoDB;
