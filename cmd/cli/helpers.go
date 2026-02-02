package main

import (
	"log"
	"os"
	"path/filepath"
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
