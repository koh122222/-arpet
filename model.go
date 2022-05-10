package main

type GetNews struct {
	NewsId int `json:"id"`
}

type DeleteNews struct {
	NewsId int `json:"id"`
}
/*
type Rubrick struct {
	RubrickId int `json:"id"`
	name string `json:"name"`
}*/

type Content struct {
	Tag string `json:"tag"`
	Text string `json:"text"`
}

type CreateNews struct {
	AvaImg string `json:"avaImg"`
	Title string `json:"title"`
	Rubrick string `json:"rubrick"`
	Content_ []Content `json:"content"'`
}