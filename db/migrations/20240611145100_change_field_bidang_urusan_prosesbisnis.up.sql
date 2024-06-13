ALTER TABLE proses_bisnis 
DROP COLUMN bidang_urusan,
ADD COLUMN bidang_urusan_id INT,
ADD index(bidang_urusan_id);