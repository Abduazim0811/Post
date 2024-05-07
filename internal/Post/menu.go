package post

import (
	"database/sql"
	"fmt"
	"os"
)

func Menu(db *sql.DB) {
	var (
		num int
	)
	fmt.Println(`
[1]Show all post
[2]My post
[3]Add like
[4]Add comment
[5]Exit
	`)
	fmt.Scanln(&num)
	switch num {
	case 1:
		Post(db)
	case 2:
		Mypost(db)
	case 3:
		Like(db)
	case 4:
		Comment(db)
	case 5:
		os.Exit(0)
	}
}
