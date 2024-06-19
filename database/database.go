package database

import (
	"database/sql"
	"log"
)

func createTables(db *sql.DB) {
	_, err := db.Exec(`
    create table anime (id integer not null primary key,title text, image text, link text);
  `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
    create table episode (id integer not null primary key,title text, src text, image text, link text, anime_id integer, foreign key(anime_id) references anime);
  `)
	if err != nil {
		log.Fatal(err)
	}
}

func NewConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "../anime.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)

	return db
}
