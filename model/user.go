package model

type User struct {
	Id       int64   `json:"id"`
	Username string  `json:"username"`
	Hash     string  `json:"hash"`
	Profile  Profile `json:"profile"`
	Created  int     `json:"created"`
}
