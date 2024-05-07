package models

// import "github.com/google/uuid"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Posts    []Post
}
