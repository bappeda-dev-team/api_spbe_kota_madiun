ALTER TABLE proses_bisnis
DROP COLUMN strategic_id,
DROP COLUMN tactical_id,
DROP COLUMN operational_id,
ADD rab_level_4_id INT,
ADD rab_level_5_id INT,
ADD rab_level_6_id INT,
ADD CONSTRAINT fk_rab_level_4 FOREIGN KEY (rab_level_4_id) REFERENCES pohon_kinerja (id),
ADD CONSTRAINT fk_rab_level_5 FOREIGN KEY (rab_level_5_id) REFERENCES pohon_kinerja (id),
ADD CONSTRAINT fk_rab_level_6 FOREIGN KEY (rab_level_6_id) REFERENCES pohon_kinerja (id);