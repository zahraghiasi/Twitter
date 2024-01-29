package models

type GeneralResponse struct {
	Ok          bool
	Result      interface{}
	Description string
	Code        int
	Error       string
}
