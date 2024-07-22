ALTER TABLE proses_bisnis 
DROP CONSTRAINT fk_operational_id,
ADD CONSTRAINT fk_operational_prosbis FOREIGN KEY (operational_id) REFERENCES pohon_kinerja (id),
ADD UNIQUE (operational_id);