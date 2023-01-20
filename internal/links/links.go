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

func GetAll() []Link {
	rows, err := database.Db.Query("SELECT id, title, address FROM Links")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var links []Link

	for rows.Next() {
		var link Link
		if err := rows.Scan(&link.ID, &link.Title, &link.Address); err != nil {
            log.Fatal(err)
        }
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

	return links
}

//Não temos ainda o User, então se fizermos um GetAll com SELECT *, ele vai esperar
//4 colunas no rows.Scan
//		sql: expected 4 destination arguments in Scan, not 3
//Como ainda não temos User, não podemos fazer nada nesse momento
//Para evitar esse erro, especificamos as colunas no SELECT