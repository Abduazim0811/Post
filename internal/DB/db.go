package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	us "Post/internal/Users"
	md "Post/internal/models"
)

func DB() {
	var (
		num int
		mp map[string]string=make(map[string]string)
		user md.User
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Abdu0811", "project")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbnames := []string{"createuser", "createpost", "createLike","createcomment"}
	for _, dbname := range dbnames {
		name := "../internal/DB/" + dbname + ".sql"
		sqlfile, err := os.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(string(sqlfile))
		if err != nil { 
			fmt.Println(dbname)
			log.Fatal("Error: ",err)
		}
	}
	name:="../internal/DB/selectuser.sql"
	sqlfile,err:=os.ReadFile(name)
	if err!=nil{
		log.Fatal("Error",err)
	}
	rows,err:=db.Query(string(sqlfile))
	if err!=nil{
		log.Fatal(err)
	}

	for rows.Next(){
		err:=rows.Scan(&user.Email,&user.Password)
		if err!=nil{
			log.Fatal(err)
		}
		mp[user.Email]=user.Password

	}
	
	fmt.Println(`
[1]Signup
[2]Signin
	`)
	fmt.Scanln(&num)
	switch num{
	case 1:
		us.Signup(db)
	case 2:
		us.Signin(db,mp)
	}

}
