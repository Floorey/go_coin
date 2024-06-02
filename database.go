package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	createTableSQL := `CREATE TABLE IF NOT EXISTS blocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		data TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveBlock(block Block) error {
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO blocks (data) VALUES (?)", string(data))
	return err

}
func LoadBlocks() ([]Block, error) {
	rows, err := db.Query("SELECT data FROM blocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blocks []Block
	for rows.Next() {
		var data string
		if err := rows.Scan(&data); err != nil {
			return nil, err
		}
		var block Block
		if err := json.Unmarshal([]byte(data), &block); err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}
