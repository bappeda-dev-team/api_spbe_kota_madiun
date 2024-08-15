package web

type UserApiData struct {
	Data []UserApiResponse `json:"data"`
}

type UserApiResponse struct {
	Nip      string `json:"nip"`
	Nama     string `json:"nama"`
	Jabatan  string `json:"jabatan"`
	UnitID   string `json:"unit_id"`
	Password string `json:"password"`
}
