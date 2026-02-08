package models

import (
	"database/sql"
	"encoding/xml"
)

type DatafileModel struct {
	DB *sql.DB
}

func (m *DatafileModel) InsertGame(game Game) error {
	stmt := `INSERT INTO games(id, name, cloneofid, description) VALUES (?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, game.ID, game.Name, game.Cloneofid, game.Description)
	if err != nil {
		return err
	}

	return nil
}

func (m *DatafileModel) InsertRom(rom Rom, gameID string) error {
	stmt := `INSERT INTO roms(game_id, text, name, size, crc, md5, sha1, sha256, status, serial, mia)
			VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt,
		gameID,
		rom.Text,
		rom.Name,
		rom.Size,
		rom.Crc,
		rom.Md5,
		rom.Sha1,
		rom.Sha256,
		rom.Status,
		rom.Serial,
		rom.Mia,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *DatafileModel) CreateDatabase() error {
	headerTable := `CREATE TABLE IF NOT EXISTS header (
	"id" INTEGER,
	"name" TEXT,
	"description" TEXT,
	"version" TEXT,
	"author" TEXT,
	"homepage" TEXT,
	"url" TEXT
	);`

	_, err := m.DB.Exec(headerTable)
	if err != nil {
		return err
	}

	gameTable := `CREATE TABLE IF NOT EXISTS games (
	"id" INTEGER PRIMARY KEY,
	"name" TEXT,
	"cloneofid" INTEGER,
	"description" TEXT);`

	_, err = m.DB.Exec(gameTable)
	if err != nil {
		return err
	}

	romTable := `CREATE TABLE IF NOT EXISTS roms (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"game_id" INTEGER,
	"text" TEXT,
	"name" TEXT,
	"size" INTEGER,
	"crc" TEXT,
	"md5" TEXT,
	"sha1" TEXT,
	"sha256" TEXT,
	"status" TEXT,
	"serial" TEXT,
	"mia" TEXT);`

	_, err = m.DB.Exec(romTable)
	if err != nil {
		return err
	}

	return nil
}

type Header struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Version     string `xml:"version"`
	Author      string `xml:"author"`
	Homepage    string `xml:"homepage"`
	URL         string `xml:"url"`
	Clrmamepro  struct {
		Text        string `xml:",chardata"`
		Forcenodump string `xml:"forcenodump,attr"`
	} `xml:"clrmamepro"`
}

type Game struct {
	Text        string   `xml:",chardata"`
	Name        string   `xml:"name,attr"`
	ID          string   `xml:"id,attr"`
	Cloneofid   string   `xml:"cloneofid,attr"`
	Description string   `xml:"description"`
	Rom         Rom      `xml:"rom"`
	Category    []string `xml:"category"`
}

type Rom struct {
	Text   string `xml:",chardata"`
	Name   string `xml:"name,attr"`
	Size   string `xml:"size,attr"`
	Crc    string `xml:"crc,attr"`
	Md5    string `xml:"md5,attr"`
	Sha1   string `xml:"sha1,attr"`
	Sha256 string `xml:"sha256,attr"`
	Status string `xml:"status,attr"`
	Serial string `xml:"serial,attr"`
	Mia    string `xml:"mia,attr"`
}

type Datafile struct {
	XMLName        xml.Name `xml:"datafile"`
	Text           string   `xml:",chardata"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Header         Header   `xml:"header"`
	Game           []Game   `xml:"game"`
}
