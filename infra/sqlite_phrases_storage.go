package infra

import (
	"database/sql"
)

type SqlitePhrasesStorage struct {
	db *sql.DB
}

func (self *SqlitePhrasesStorage) SavePhrase(phrase string) {
	_, err := self.db.Exec("INSERT INTO phrases VALUES (?)", phrase)
	if err != nil {
		panic(err)
	}
}

func (self *SqlitePhrasesStorage) GetPhrase(target_substr string) string {
	rows, err := self.db.Query("SELECT * FROM phrases WHERE text LIKE '%' || $1 || '%'", target_substr)
	if err != nil {
		panic(err)
	}

	if !rows.Next() {
		panic("No text")
	}

	var text string
	rows.Scan(&text)

	return text
}

func CreateSqlitePhrasesStorage(db *sql.DB) *SqlitePhrasesStorage {
	return &SqlitePhrasesStorage{
		db: db,
	}
}
