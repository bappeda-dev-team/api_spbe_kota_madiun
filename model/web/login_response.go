package web

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	NIP     string `json:"nip"`
	Nama    string `json:"nama"`
	KodeOPD string `json:"kode_opd"`
	Roles   []Role `json:"roles"`
}

type Role struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
