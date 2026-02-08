package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/w0/sqlomatic/internal/models"
	_ "modernc.org/sqlite"
)

type application struct {
	logger    *slog.Logger
	dataflies *models.DatafileModel
}

func main() {
	datPath := flag.String("datafile", "", "Path to the datafile to convert.")
	flag.Parse()

	if *datPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	resolved, err := resolvePath(*datPath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	filename := filepath.Base(resolved)

	dat, err := readFile(resolved)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	dsn := fmt.Sprintf("file:%s.db?cached=shared", filename[:len(filename)-4])

	db, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := application{
		logger:    logger,
		dataflies: &models.DatafileModel{DB: db},
	}

	err = app.dataflies.CreateDatabase()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	for _, game := range dat.Game {
		err = app.dataflies.InsertGame(game)
		if err != nil {
			logger.Warn(err.Error())
		}

		err = app.dataflies.InsertRom(game.Rom, game.ID)
		if err != nil {
			logger.Warn(err.Error())
		}
	}

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
