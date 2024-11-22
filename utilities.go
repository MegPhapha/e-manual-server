package main

import (
	"os"
	"errors"
	"strings"
)

func scandir(dirPath string) ([]string, error) {
	f, err := os.Open(dirPath)
	if err != nil{
		return []string{}, errors.New("unable to access directory" + dirPath)
	}

	files, err := f.Readdir(0)
	if err != nil {
		return []string{}, errors.New("unable to access directory " + dirPath)
	}
	result := make ([]string, len(files))
	for i := range files {
		result[i] = strings.Split(files[i].Name(), ".")[0]
	}
	return result, nil
}