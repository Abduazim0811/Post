package models


import "time"

type Post struct {
	Id         uint
	User_id    int
	Title      string
	Content    string
	Created_at time.Time
}