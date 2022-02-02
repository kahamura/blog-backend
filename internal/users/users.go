package users

import (
	"database/sql"
	"log"
	"time"

	database "github.com/ka-aki/blog-backend/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"name"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO users (username, password, created_at, updated_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, hashedPassword, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("select ID from users where username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return Id, nil
}

func (user *User) Authentificate() bool {
	statement, err := database.Db.Prepare("select Password from users where username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Username)

	var hasedPassword string
	err = row.Scan(&hasedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hasedPassword)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
