package response

import "github.com/ghiac/social/internal/repositories"

type LastInsertIDResponse struct {
	ID uint64 `json:"ID" xml:"ID" form:"ID" query:"ID"`
}

type GeneralResponse struct {
	Ok          bool
	Result      interface{}
	Description string
	Code        int
	Error       string
}

type LoginResponse struct {
	Token string
	User  repositories.User
}
