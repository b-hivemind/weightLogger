package db

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPass := h.Sum(nil)
	return fmt.Sprintf("#%x@", hashedPass)
}

func FindUser(user User) (User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username='%s'", usersTable, user.Username)
	db, err := connect()
	if err != nil {
		return User{}, fmt.Errorf("FindUser | Error connecting to database | %v", err)
	}
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return User{}, fmt.Errorf("FindUser: %v", err)
	}
	defer rows.Close()
	tempUser := User{}
	for rows.Next() {
		err = rows.Scan(&tempUser.UUID, &tempUser.Username, &tempUser.Password)
		if err != nil {
			return User{}, fmt.Errorf("FindUser: %v", err)
		}
	}
	if tempUser.Password == "" {
		return tempUser, nil
	}
	if HashPassword(user.Password) != tempUser.Password {
		return User{Password: "-1"}, nil
	}
	return tempUser, nil
}

func RegisterUser(user User) (User, error) {
	newUser := User{
		UUID:     uuid.New(),
		Username: user.Username,
		Password: HashPassword(user.Password),
	}
	query := fmt.Sprintf("INSERT INTO %s (uid, username, password) VALUES ('%s', '%s', '%s')", usersTable, newUser.UUID, newUser.Username, newUser.Password)
	db, err := connect()
	if err != nil {
		return newUser, fmt.Errorf("RegisterUser | Error connecting to database | %v", err)
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		return newUser, fmt.Errorf("RegisterUser %v", err)
	}
	return newUser, nil
}
