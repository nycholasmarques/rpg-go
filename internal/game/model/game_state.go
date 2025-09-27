package model

type GameState struct {
	Character     Character `json:"character"`
	PosX          float64   `json:"pos_x"`
	PosY          float64   `json:"pos_y"`
	Filename_save string    `json:"filename_save"`
	DebugMsg      string
}
