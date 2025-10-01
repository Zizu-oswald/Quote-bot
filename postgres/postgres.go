package postgres

import (
	"database/sql"
	// "fmt"

	"github.com/Zizu-oswald/Quote-bot/telegram"

	_ "github.com/lib/pq"
)

// var table = `create table users (chatid integer, lang text);`
// insert into users (chatid, lang) values (1, 'ru'), (2, 'en');
type Database *sql.DB

func ConnectToSql() (*sql.DB, error) {
	connStr := "user=myuser password=mysecretpassword dbname=mydatabase sslmode=disable"
	var db Database 
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// result, err := db.Query("select * from users;")
	// if err != nil {
	// 	return nil, err
	// }

	// var mass []telegram.ChatStruct

	// for result.Next() {
	// 	c := telegram.ChatStruct{}
	// 	err := result.Scan(&c.ID, &c.Lang, &c.LastMessageID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	mass = append(mass, c)
	// }

	// fmt.Println(mass)
	return db, nil
}

func AddUser(db *sql.DB, u telegram.ChatStruct) error {
	_, err := db.Exec("insert into users (chatid, lang, lastmessageid) values ($1, $2, $3);", u.ID, u.Lang, u.LastMessageID)
	return err
}
