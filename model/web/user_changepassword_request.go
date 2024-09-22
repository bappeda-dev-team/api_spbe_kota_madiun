package web

type ChangePasswordRequest struct {
	UserID      int    `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
