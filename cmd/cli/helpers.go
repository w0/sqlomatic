package main

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"

	"github.com/sqlomatic/internal/models"
)

func resolvePath(datafile string) string {

	abs, err := filepath.Abs(datafile)
	if err != nil {
		log.Fatalln(err.Error())
	}

	info, err := os.Stat(abs)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if info.IsDir() {
		log.Fatalln("error: path is a directory not a datafile.")
	}

	return abs
}

func readFile(path string) models.Datafile {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var dat models.Datafile

	decoder := xml.NewDecoder(file)

	err = decoder.Decode(&dat)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return dat

}
