package users

import (
	"database/sql"
	"fmt"
	md "Post/internal/models"
	ps "Post/internal/Post"
)

func Signin(db *sql.DB, mp map[string]string) {
	var (
		user md.User
	)

	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Password: ")
	fmt.Scanln(&user.Password)

	value, found := mp[user.Email]
	if found {
		if user.Password == value {
			fmt.Println("Siz muvaffaqiyatli kirdingiz!!")
			ps.Menu(db)

		} else {
			fmt.Println("Parol xato!!")
		}
	} else {
		fmt.Println("Email xato!!")
	}

}
