package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
	//fmt.Println(r.Header.Get("Content-Type")) TODO можно добавить этот тип
	if r.Method == "GET" {
		var deleteNews DeleteNews
		var unmarshalErr *json.UnmarshalTypeError
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&deleteNews)
		if err != nil {
			if errors.As(err, &unmarshalErr) {
				errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
			} else {
				errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
			}
			return
		}

		db_global.GetContext(context.TODO(), &result, `select * FROM news.news_get_news_with_id($1)`,
			deleteNews.newsId)
		fmt.Println(deleteNews.newsId)
		fmt.Println(r.Body)
	} else if r.Method == "DELETE" {
		result_delete := ""
		db_global.GetContext(context.TODO(), &result_delete, `select * FROM news.news_delete_news(4)`)
		if result_delete == "true" {
			result = "true"
		} else {
			result = "false"
		}
	}

	fmt.Fprint(w, result)


}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
