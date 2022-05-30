package model

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Profile  Profile `json:"profile"`
	Created  int     `json:"created"`
}
