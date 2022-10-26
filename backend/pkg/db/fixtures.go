package db

import "github.com/google/uuid"

const (
	weightLogTable = "main"
	usersTable     = "users"
)

type Entry struct {
	Date   int64 `json:"date"`
	UID    string
	Weight float32 `json:"weight"`
}

type User struct {
	UUID     uuid.UUID `json:"uid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
