package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	md "Post/internal/models"
	ps "Post/internal/Post"
)

func Signup(db *sql.DB) {
	var (
		user md.User
	)

	fmt.Print("Name: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Password: ")
	fmt.Scanln(&user.Password)

	name := "../internal/DB/insertusers.sql"

	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlfile), user.Name,user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----Siz muvvaqiyatli kirdingiz!!!----")
	ps.Menu(db)

}
