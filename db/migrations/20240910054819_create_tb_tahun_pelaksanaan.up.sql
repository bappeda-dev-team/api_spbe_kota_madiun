CREATE TABLE tahun_pelaksanaan(
    id INT NOT NULL AUTO_INCREMENT,
    id_rencana_pelaksana INT,
    tahun INT,
    PRIMARY KEY (id),
    CONSTRAINT fk_id_rencana_pelaksana FOREIGN KEY (id_rencana_pelaksana) REFERENCES rencana_pelaksanaan (id)
)ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;