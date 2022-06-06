package model

type Travel struct {
	Id     int    `json:"id"`
	Driver Driver `json:"driver"`
	Start  int    `json:"start"`
	End    *int   `json:"end"`
}
