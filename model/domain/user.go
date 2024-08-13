package domain

type User struct {
	ID       int
	NIP      string
	Nama     string
	KodeOPD  string
	Password string
	Roles    []Role
}

type Role struct {
	ID   int
	Nama string
}
