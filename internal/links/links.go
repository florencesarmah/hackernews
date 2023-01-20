package links

import(
	database "github.com/florencesarmah/hackernews/internal/pkg/db/migrations/mysql"
	"github.com/florencesarmah/hackernews/internal/users"
	"log"
)

type Link struct {
	ID		string
	Title	string
	Address	string
	User	*users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Row inserted!")
	return id
}