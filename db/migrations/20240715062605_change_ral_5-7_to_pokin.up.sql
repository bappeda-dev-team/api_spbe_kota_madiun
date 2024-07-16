ALTER TABLE layanan_spbe DROP FOREIGN KEY fk_ral_level_5;
ALTER TABLE layanan_spbe DROP COLUMN ral_level_5_id;
ALTER TABLE layanan_spbe ADD strategic_id INT;
ALTER TABLE layanan_spbe ADD CONSTRAINT fk_startegic_layanan FOREIGN KEY (strategic_id) REFERENCES pohon_kinerja (id);
ALTER TABLE layanan_spbe ADD tactical_id INT;
ALTER TABLE layanan_spbe ADD operational_id INT;
ALTER TABLE layanan_spbe ADD CONSTRAINT fk_tactical_layanan FOREIGN KEY (tactical_id) REFERENCES pohon_kinerja (id);
ALTER TABLE layanan_spbe ADD CONSTRAINT fk_operational_layanan FOREIGN KEY (operational_id) REFERENCES pohon_kinerja (id);