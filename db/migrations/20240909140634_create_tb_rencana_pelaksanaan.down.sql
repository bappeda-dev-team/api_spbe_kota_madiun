ALTER TABLE rencana_pelaksanaan
DROP FOREIGN KEY fk_kebutuhan_spbe,
DROP FOREIGN KEY fk_sasarankinerja_pegawai,
DROP FOREIGN KEY fk_perangkatdaerah,
DROP FOREIGN KEY fk_kode_opd_rencana;

DROP TABLE IF EXISTS rencana_pelaksanaan;
