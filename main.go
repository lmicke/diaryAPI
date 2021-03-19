package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/lmicke/diaryAPI/handlers"
)

func main() {

	db, err := sql.Open(
		"mysql",
		"api:1234@tcp(127.0.0.1:3306)/diary",
	)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/entry/{id}", handlers.MakeGetEntry(db)).Methods("GET")
	r.Handle("/entry", handlers.MakeCreateEntry(db)).Methods("POST")
	r.Handle("/entry/{id}", handlers.MakeDeleteEntry(db)).Methods("DELETE")
	//http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
