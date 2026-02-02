package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sqlomatic/internal/models"
	_ "modernc.org/sqlite"
)

type application struct {
	dataflies *models.DatafileModel
}

func main() {
	datPath := flag.String("datafile", "", "Path to the datafile to convert.")
	flag.Parse()

	if *datPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	resolved := resolvePath(*datPath)
	filename := filepath.Base(resolved)

	dsn := fmt.Sprintf("file:%s.db?cached=shared", filename[:len(filename)-4])

	log.Println(dsn)

	db, err := openDB(dsn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer db.Close()

	app := application{
		dataflies: &models.DatafileModel{DB: db},
	}

	log.Println(app)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
