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

func GetAll() []Article {
	stmt, err := database.Db.Prepare("select id, title, content, created_at, updated_at from articles")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return articles
}

func (article Article) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO articles(title, content, created_at, updated_at) VALUES(?,?,?,?)")
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
