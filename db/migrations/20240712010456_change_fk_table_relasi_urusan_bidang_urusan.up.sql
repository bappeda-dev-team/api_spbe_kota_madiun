ALTER TABLE urusan_bidang_opd DROP FOREIGN KEY fk_bidang_urusan;
ALTER TABLE urusan_bidang_opd ADD CONSTRAINT fk_bidang_urusan FOREIGN KEY (bidang_urusan) REFERENCES bidang_urusan (bidang_urusan);