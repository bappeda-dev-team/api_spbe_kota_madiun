ALTER TABLE proses_bisnis ADD CONSTRAINT fk_sasaran_kota FOREIGN KEY (sasaran_kota) REFERENCES sasaran_kota (sasaran);