package main

import(
	db "Post/internal/DB"
	_ "github.com/lib/pq"
)

func main(){
	db.DB()
}