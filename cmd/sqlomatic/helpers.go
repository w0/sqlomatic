package main

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"github.com/w0/sqlomatic/internal/models"
)

func resolvePath(datafile string) (string, error) {

	abs, err := filepath.Abs(datafile)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(abs)
	if err != nil {
		return "", err
	}

	return abs, nil
}

func readFile(path string) (models.Datafile, error) {

	file, err := os.Open(path)
	if err != nil {
		return models.Datafile{}, err
	}

	var dat models.Datafile

	decoder := xml.NewDecoder(file)

	err = decoder.Decode(&dat)

	if err != nil {
		return models.Datafile{}, err
	}

	return dat, nil

}
