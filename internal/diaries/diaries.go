package diaries

import (
	"log"
	"time"

	database "github.com/ka-aki/blog-backend/internal/pkg/db/mysql"
	"github.com/ka-aki/blog-backend/internal/users"
)

type Diary struct {
	ID        string
	Title     string
	Content   string
	User      *users.User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAll() []Diary {
	stmt, err := database.Db.Prepare("select id, title, content, created_at, updated_at from diaries")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var diaries []Diary
	for rows.Next() {
		var diary Diary
		err := rows.Scan(&diary.ID, &diary.Title, &diary.Content, &diary.CreatedAt, &diary.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		diaries = append(diaries, diary)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return diaries
}

func (diary Diary) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO diaries(title, content, created_at, updated_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(diary.Title, diary.Content, diary.CreatedAt, diary.UpdatedAt)
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
