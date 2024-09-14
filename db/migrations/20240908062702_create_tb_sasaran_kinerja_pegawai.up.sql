CREATE TABLE sasaran_kinerja_pegawai(
    id int NOT NULL AUTO_INCREMENT,
    kode_opd VARCHAR(255),
    kode_sasaran VARCHAR(255),
    tahun_sasaran VARCHAR(255),
    sasaran_kinerja VARCHAR(255),
    anggaran_sasaran VARCHAR(255),
    pelaksana_sasaran VARCHAR(255),
    kode_subkegiatan_sasaran VARCHAR(255),
    subkegiatan_sasaran VARCHAR(255),
    PRIMARY KEY (id)
)ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;