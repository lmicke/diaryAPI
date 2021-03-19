package model

import (
	"database/sql"
	"log"
)

//Entry is the general Entry struct.
type Entry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Date  string `json:"date"`
}

func (e *Entry) GetEntry(id int, db *sql.DB) {
	row := db.QueryRow("SELECT id, title, text, created_at FROM entries WHERE id=?;", id)
	err := row.Scan(&e.ID, &e.Title, &e.Text, &e.Date)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n")
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("id is %v, Title is %s, Text is %s\n", e.ID, e.Title, e.Text)

	}
}

func (e *Entry) CreateEntry(db *sql.DB) {
	db.QueryRow("INSERT INTO entries (title,text) VALUES(?, ?);", e.Title, e.Text)
	row := db.QueryRow("SELECT id, title, text, created_at FROM entries WHERE title=? AND text=?;", e.Title, e.Text)
	err := row.Scan(&e.ID, &e.Title, &e.Text, &e.Date)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", e.ID)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("id is %v, Title is %s, Text is %s\n", e.ID, e.Title, e.Text)

	}
}

func (e *Entry) DeleteEntry(db *sql.DB) {
	db.QueryRow("DELETE FROM entries WHERE id=?;", e.ID)
	row := db.QueryRow("SELECT id, title, text, created_at FROM entries WHERE id=?;", e.ID)
	err := row.Scan(&e.ID, &e.Title, &e.Text, &e.Date)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", e.ID)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("id is %v, Title is %s, Text is %s\n", e.ID, e.Title, e.Text)
	}
}

func (e *Entry) UpdateEntry(db *sql.DB) {
	if e.Exists(db) {
		db.QueryRow(`UPDATE entries SET text=? , title=? WHERE id=?;`, e.Text, e.Title, e.ID)
		row := db.QueryRow("SELECT id, title, text, created_at FROM entries WHERE id=?;", e.ID)
		err := row.Scan(&e.ID, &e.Title, &e.Text, &e.Date)

		switch {
		case err == sql.ErrNoRows:
			log.Printf("no user with id %d\n", e.ID)
		case err != nil:
			log.Fatalf("query error: %v\n", err)
		default:
			log.Printf("id is %v, Title is %s, Text is %s\n", e.ID, e.Title, e.Text)
		}
	} else {
		e.CreateEntry(db)
	}

}

func (e *Entry) Exists(db *sql.DB) bool {
	row := db.QueryRow("SELECT id FROM entries WHERE id=?;", e.ID)
	err := row.Scan(&e.ID)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		return true

	}
	return true
}

/*
Base Database;
MariaDB [diary]> Describe entries;
+------------+--------------+------+-----+---------------------+----------------+
| Field      | Type         | Null | Key | Default             | Extra          |
+------------+--------------+------+-----+---------------------+----------------+
| id         | int(11)      | NO   | PRI | NULL                | auto_increment |
| title      | varchar(255) | NO   |     | NULL                |                |
| date       | date         | YES  |     | NULL                |                |
| text       | text         | YES  |     | NULL                |                |
| created_at | timestamp    | NO   |     | current_timestamp() |                |
+------------+--------------+------+-----+---------------------+----------------+

*/

/*
CREATE TABLE IF NOT EXISTS entries ( id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) NOT NULL, text TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP );
INSERT INTO entries (title,text) VALUES("Test", "Hallo Welt");

SELECT * FROM entries;
+----+-------+------------+---------------------+
| id | title | text       | created_at          |
+----+-------+------------+---------------------+
|  1 | Test  | Hallo Welt | 2021-03-14 17:53:10 |
|  2 | Test  | Hallo Welt | 2021-03-14 17:54:13 |
+----+-------+------------+---------------------+


SELECT *  FROM entries WHERE id=1;
+----+-------+------------+---------------------+
| id | title | text       | created_at          |
+----+-------+------------+---------------------+
|  1 | Test  | Hallo Welt | 2021-03-14 17:53:10 |
+----+-------+------------+---------------------+

DELETE FROM entries WHERE id=1;

UPDATE entries SET text="Mein erster Eintrag!" WHERE id=1;


*/
