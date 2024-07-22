ALTER TABLE aplikasi 
ADD COLUMN raa_level_4_id INT,
ADD CONSTRAINT fk_raa_level_4 FOREIGN KEY (raa_level_4_id) REFERENCES referensi_arsitekturs (id);