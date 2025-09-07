package game

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/nycholasmarques/rpg-go/internal/game/model"
)

const path string = "saves"

var (
	ErrToRemoveFile = errors.New("error to remove save file")
)

func Save(gamestate model.GameState, filename string) {
	data, err := json.Marshal(gamestate)
	if err != nil {
		panic(err)
	}
	
	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	if gamestate.Filename_save == "" && filename != "" {
		err = os.WriteFile(path + "/" + filename, data, 0644)
		if err != nil {
			panic(err)
		}
	}
	err = os.WriteFile(path + "/" + gamestate.Filename_save, data, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadSave(filename string) model.GameState {
	read, err := os.ReadFile(path + "/" + filename)
	if err != nil {
		panic(err)
	}

	var gameState model.GameState
	err = json.Unmarshal(read, &gameState)
	if err != nil {
		panic(err)
	}

	return gameState
}

func DeleteSave(filename string) error {
	fileInfo, err := os.Stat(path + "/" + filename) 
	if errors.Is(err, os.ErrNotExist) {
		return os.ErrNotExist
	}
	err = os.Remove(path + "/" + fileInfo.Name())
	if err != nil {
		return ErrToRemoveFile
	}
	return nil
}

func PrintSaves() []string {
	files, err := os.ReadDir("saves")
	if err != nil {
		panic(err)
	}
	var filesForArray []string
	for _, v := range files {
		filesForArray = append(filesForArray, v.Name())
	}
	return filesForArray
}

