package models

import (
	"database/sql"
	"encoding/xml"
)

type DatafileModel struct {
	DB *sql.DB
}

type Datafile struct {
	XMLName        xml.Name `xml:"datafile"`
	Text           string   `xml:",chardata"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Header         struct {
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
	} `xml:"header"`
	Game []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		ID          string `xml:"id,attr"`
		Cloneofid   string `xml:"cloneofid,attr"`
		Description string `xml:"description"`
		Rom         struct {
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
		} `xml:"rom"`
		Category []string `xml:"category"`
	} `xml:"game"`
}
