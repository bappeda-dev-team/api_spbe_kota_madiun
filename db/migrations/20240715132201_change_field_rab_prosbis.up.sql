ALTER TABLE proses_bisnis 
DROP FOREIGN KEY fk_rab_level_4,
DROP FOREIGN KEY fk_rab_level_5,
DROP FOREIGN KEY fk_rab_level_6,
DROP COLUMN rab_level_4_id,
DROP COLUMN rab_level_5_id,
DROP COLUMN rab_level_6_id,
ADD strategic_id INT,
ADD tactical_id INT,
ADD operational_id INT;