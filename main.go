package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
)

var db_global *sqlx.DB

func rubricksHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//response := fmt.Sprintf("Product %s", id)

	//fmt.Fprint(w, r.Method)
	result := "error get rubricks?"
	db_global.GetContext(context.TODO(), &result, `select * FROM news.news_get_rubricks()`)
	fmt.Fprint(w, result)
}

func postIdHandler(w http.ResponseWriter, r *http.Request) {
	result := "error get news?"
	if (r.Method == "GET") {
		delete_news := DeleteNews{}
		json.NewDecoder(r.Body).Decode(&delete_news)
		db_global.GetContext(context.TODO(), &result, `select * FROM news.news_get_news_with_id($1)`,
			delete_news.id)
		fmt.Println(delete_news.id)
		fmt.Println(r.Body)
	} else if (r.Method == "DELETE") {
		result_delete := ""
		db_global.GetContext(context.TODO(), &result_delete, `select * FROM news.news_delete_news(4)`)
		if (result_delete == "true") {
			result = "true"
		} else {
			result = "false"
		}
	}

	fmt.Fprint(w, result)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rubricks", rubricksHandler)
	router.HandleFunc("/post/id", postIdHandler)

	http.Handle("/",router)

	//router = mux.NewRouter()
	//router.HandleFunc("/post/id", postIdHandler)
	//http.Handle("/", router)

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