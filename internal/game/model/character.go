package model

type levels int

const (
	Level_1 levels = 0
	Level_2 levels = 100
)

type Character struct {
	Name  string `json:"name"`
	Level levels `json:"level"`
	Hp    int    `json:"hp"`
	Xp    int    `json:"xp"`
	Class Class  `json:"class"`
}
