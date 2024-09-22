package domain

import "database/sql"

type User struct {
	ID       int
	NIP      string
	Nama     string
	KodeOPD  string
	Jabatan  sql.NullString // Ubah ini dari string ke sql.NullString
	Password string
	Roles    []Role
}

type Role struct {
	ID   int
	Nama string
}
