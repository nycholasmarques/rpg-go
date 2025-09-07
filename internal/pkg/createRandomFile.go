package pkg

import (
	"os"
	"strconv"
)

func CreateRandomFile(localPath string, filename string) (string, error) {
	files, err := os.ReadDir(localPath)
	if err != nil {
		panic(err)
	}

	for _, v := range files {
		if v.Name() == filename {
			newFileName := filename + strconv.Itoa(len(files) + 1)
			return newFileName, nil
		}
	}

	return filename, nil
}