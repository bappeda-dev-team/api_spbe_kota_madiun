ALTER TABLE proses_bisnis
ADD CONSTRAINT fk_strategic FOREIGN KEY (strategic_id) REFERENCES pohon_kinerja(id),
ADD CONSTRAINT fk_tactical FOREIGN KEY (tactical_id) REFERENCES pohon_kinerja(id),
ADD CONSTRAINT fk_operational_id FOREIGN KEY (operational_id) REFERENCES pohon_kinerja(id);