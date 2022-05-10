package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
)

var db_global *sqlx.DB

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rubricks", rubricksHandler)
	router.HandleFunc("/post/id", postIdHandler)
	router.HandleFunc("/post", createPostHandler)
	router.HandleFunc("/posts", getArrayPostHandler)

	http.Handle("/",router)

	connStr := "user=postgres password=zybrjcnz dbname=siteDB sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db_global = db

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}