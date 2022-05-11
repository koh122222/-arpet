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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	result := "error get rubricks?"
	db_global.GetContext(context.TODO(), &result, `select * FROM news.news_get_rubricks()`)
	fmt.Fprint(w, result)
}

func postIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	result := "error get news?"
	//fmt.Println(r.Header.Get("Content-Type")) TODO можно добавить этот тип
	if r.Method == "GET" {
		var getNews GetNews
		var unmarshalErr *json.UnmarshalTypeError
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&getNews)
		if err != nil {
			if errors.As(err, &unmarshalErr) {
				errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
			} else {
				errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
			}
			return
		}

		db_global.GetContext(context.TODO(), &result, `select * FROM news.news_get_news_with_id($1)`,
			getNews.NewsId)
	} else if r.Method == "DELETE" {
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
		result_delete := ""
		db_global.GetContext(context.TODO(), &result_delete, `select * FROM news.news_delete_news($1)`,
			deleteNews.NewsId)
		if result_delete == "true" {
			result = "true"
		} else {
			result = "false"
		}
	} else if r.Method == "PUT" {

	}

	fmt.Fprint(w, result)


}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	result := "error create post?"
	var createNews CreateNews
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&createNews)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	fmt.Println(createNews)
	//TODO доделать создание новости
	/*
	result_delete := ""

	db_global.GetContext(context.TODO(), &result_delete, `select * FROM news.news_delete_news($1)`,
		deleteNews.NewsId)*/

	fmt.Fprint(w, result)
}

func getArrayPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rubrick := r.URL.Query().Get("rubrick")
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	search :=  r.URL.Query().Get("search")
	result := "error get array?"
	db_global.GetContext(context.TODO(), &result, `select * FROM news.news_get_array_news($1, $2, $3, $4)`,
		rubrick, startDate, endDate, search)
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
