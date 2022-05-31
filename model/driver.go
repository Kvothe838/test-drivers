package model

type Driver struct {
	DNI     string `json:"DNI"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	User    User   `json:"user"`
}
