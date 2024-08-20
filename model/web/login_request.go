package web

type LoginRequest struct {
	NIP      string `json:"nip"`
	Password string `json:"password"`
}
