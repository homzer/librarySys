package model

type Paper struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Theme  string `json:"theme"`
}
