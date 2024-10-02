ALTER TABLE rencana_pelaksanaan 
ADD COLUMN indikator_perangkatdaerah ENUM('internal', 'eksternal'),
ADD COLUMN perangkat_daerah VARCHAR(255),
ADD CONSTRAINT fk_perangkatdaerah FOREIGN KEY (perangkat_daerah) REFERENCES opd (kode_opd);