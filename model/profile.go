package model

type Profile struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}
