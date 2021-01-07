package links

import (
	database "hackerclone/internal/pkg/db"
	"hackerclone/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)

	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("error", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID")

	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var username string
	var id string
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)

		}

		link.User = &users.User{
			ID:       id,
			Username: username,
		}
		links = append(links, link)

	}
	if err != nil {
		log.Fatal(err)

	}
	return links
}

func (link *Link) Delete(id int) (int64, error) {

	print(id)

	res, err := database.Db.Exec("delete from Links where id = ?", id)
	print(res)
	if err != nil {
		log.Fatal(err)
	}
	print("row deleted")
	return res.RowsAffected()

}
