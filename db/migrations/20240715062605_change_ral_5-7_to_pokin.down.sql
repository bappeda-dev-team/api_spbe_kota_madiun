ALTER TABLE layanan_spbe ADD COLUMN ral_level_5_id INT;
ALTER TABLE layanan_spbe ADD CONSTRAINT fk_ral_level_5 FOREIGN KEY (ral_level_5_id) REFERENCES referensi_arsitekturs (id);
ALTER TABLE layanan_spbe DROP FOREIGN KEY fk_startegic_layanan;
ALTER TABLE layanan_spbe DROP FOREIGN KEY fk_tactical_layanan;
ALTER TABLE layanan_spbe DROP FOREIGN KEY fk_operational_layanan;
ALTER TABLE layanan_spbe DROP strategic_id;
ALTER TABLE layanan_spbe DROP tactical_id;
ALTER TABLE layanan_spbe DROP operational_id;