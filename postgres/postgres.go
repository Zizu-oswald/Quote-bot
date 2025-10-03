package postgres

import (
	"database/sql"
	"log"
	// "fmt"

	"github.com/Zizu-oswald/Quote-bot/telegram"

	_ "github.com/lib/pq"
)

// var table = `create table users (chatid integer, lang text);`
// insert into users (chatid, lang) values (1, 'ru'), (2, 'en');
type Database struct {
	Db *sql.DB
}

func (d *Database) ConnectToSql() (error) {
	connStr := "user=myuser password=mysecretpassword dbname=mydatabase sslmode=disable"
	// var db Database 
	var err error
	db, err := sql.Open("postgres", connStr)
	d.Db = db
	if err != nil {
		return err
	}
	return nil

}

func (d *Database) AddUser(u telegram.ChatStruct) error {
	_, err := d.Db.Exec("insert into users (chatid, lang, lastmessageid) values ($1, $2, $3);", u.ID, u.Lang, u.LastMessageID)
	return err
}

func (d *Database) Close(){
	err := d.Db.Close()
	if err != nil {
		log.Println("error with closing database: ", err)
	}
}

func (d *Database) TakeUser(id int) (telegram.ChatStruct, error) {
	result:= d.Db.QueryRow("select * from users where chatid = $1;", id)

	var chat telegram.ChatStruct
	err := result.Scan(&chat.ID, &chat.Lang, &chat.LastMessageID)

	if err != nil {
	  return telegram.ChatStruct{}, err
	}
	return chat, nil
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
// return db, nil