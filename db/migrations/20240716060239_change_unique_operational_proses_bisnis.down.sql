ALTER TABLE proses_bisnis 
ADD CONSTRAINT fk_operational_id FOREIGN KEY (operational_id) REFERENCES pohon_kinerja (id),
DROP CONSTRAINT fk_operational_prosbis,
DROP INDEX operational_id;