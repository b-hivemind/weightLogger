package db

import "github.com/google/uuid"

const (
	weightLogTable = "main"
	usersTable     = "users"
)

type Entry struct {
	Date   string  `json:"date"`
	Weight float32 `json:"weight"`
}

type User struct {
	UUID     uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
