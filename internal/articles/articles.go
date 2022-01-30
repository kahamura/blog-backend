package articles

import (
	"log"
	"time"

	database "github.com/ka-aki/blog-backend/internal/pkg/db/mysql"
	"github.com/ka-aki/blog-backend/internal/users"
)

type Article struct {
	ID        string
	Title     string
	Content   string
	User      *users.User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (article Article) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO articles(title,content,created_at,updated_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(article.Title, article.Content, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")

	return id
}
