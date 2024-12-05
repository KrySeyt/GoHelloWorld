package main

import (
	"database/sql"
	"net/http"
	"packname/app"
	"packname/infra"
	"packname/presentation"

	"github.com/go-resty/resty/v2"
	_ "github.com/mattn/go-sqlite3"
)

func create_db() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE TABLE phrases(text VARCHAR)")

	return db
}

func main() {
	db := create_db()
	defer db.Close()

	http_client := resty.New()

	http_phrases_dictionary := infra.CreateHttpPhrasesDictionary(http_client, "https://httpbin.org/")

	sqlite_phrases_storage := infra.CreateSqlitePhrasesStorage(db)
	transaction_manager := infra.CreateSqlTransactionManager(db)

	say_hello := app.CreateSayHello(
		sqlite_phrases_storage,
		http_phrases_dictionary,
		transaction_manager,
	)

	get_hello := app.CreateGetHello(sqlite_phrases_storage)

	api := presentation.CreateHelloHttpApi(
		say_hello,
		get_hello,
	)

	http.HandleFunc("/get-hello", api.GetHello)
	http.HandleFunc("/say-hello", api.SayHello)

	http.ListenAndServe(":8000", nil)

}
