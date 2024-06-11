-- ALTER TABLE proses_bisnis CONSTRAINT fk_sasaran_kota_id FOREIGN KEY (sasaran_kota_id) REFERENCES sasaran_kota (id);
ALTER TABLE proses_bisnis ADD FOREIGN KEY (sasaran_kota_id) REFERENCES sasaran_kota (id);
