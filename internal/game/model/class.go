package model

type Class struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Hp   int    `json:"hp"`
	Atk  int    `json:"atk"`
	Def  int    `json:"def"`
}
