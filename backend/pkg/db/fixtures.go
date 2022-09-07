package db

const (
	weightLogTable = "main"
	usersTable     = "users"
)

type Entry struct {
	Date   string  `json:"date"`
	Weight float32 `json:"weight"`
}
