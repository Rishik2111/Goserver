package models

import (
	"database/sql"
	"log"
)

func CreateTablesinDB(db *sql.DB) {
	// Create the necessary Tables(users, movies and relation b/w them) in the Database if they don't exist.
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			userid SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			favgenres VARCHAR(255),
			createdat TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS movies (
			movieid SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			genre VARCHAR(255) NOT NULL,
			storyline VARCHAR(255) NOT NULL,
			createdat TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS usermovierelation (
			userid INT NOT NULL,
			movieid INT NOT NULL,
			rating INT CHECK (rating >= 0 AND rating <= 10),
			PRIMARY KEY (userid, movieid),
			FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE,
			FOREIGN KEY (movieid) REFERENCES movies(movieid) ON DELETE CASCADE
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}
