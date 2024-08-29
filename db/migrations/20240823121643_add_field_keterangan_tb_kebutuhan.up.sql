ALTER TABLE kebutuhan_spbe 
ADD COLUMN keterangan TEXT,
ADD COLUMN indikator_pj ENUM('internal', 'eksternal'),
ADD COLUMN pj VARCHAR(255),
ADD CONSTRAINT fk_penanggung_jawab_kode_opd FOREIGN KEY (pj) REFERENCES opd (kode_opd);