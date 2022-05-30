package model

type Travel struct {
	Driver Driver `json:"driver"`
	Start  int    `json:"start"`
	End    *int   `json:"end"`
}
