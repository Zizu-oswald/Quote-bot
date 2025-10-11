package telegram

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "fmt"

	_ "github.com/lib/pq"
)

// var table = `create table users (chatid integer, lang text);`
// insert into users (chatid, lang) values (1, 'ru'), (2, 'en');
type Database struct {
	Db *sql.DB
}

func (d *Database) ConnectToSql() error {
	connectionStr := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error
	db, err := sql.Open("postgres", connectionStr)
	d.Db = db
	if err != nil {
		return err
	}
	return nil

}

func MakeTable(d *Database) error {
	_, err := d.Db.Exec("create table if not exists users ( chatid bigint, lang text, lastmessageid integer );")
	return err
}

func (d *Database) AddUser(u ChatStruct) error {
	_, err := d.Db.Exec("insert into users (chatid, lang, lastmessageid) values ($1, $2, $3);", u.ID, u.Lang, u.LastMessageID)
	return err
}

func (d *Database) Close() {
	err := d.Db.Close()
	if err != nil {
		log.Println("error with closing database: ", err)
	}
}

func (d *Database) GetUser(id int64) (ChatStruct, error) {
	result := d.Db.QueryRow("select * from users where chatid = $1;", id)

	var chat ChatStruct
	err := result.Scan(&chat.ID, &chat.Lang, &chat.LastMessageID)

	if err != nil {
		return ChatStruct{}, err
	}
	return chat, nil
}

func (d *Database) UpdateUserData(u ChatStruct) error {
	_, err := d.Db.Exec("update users set lang = $1, lastmessageid = $2 where chatid = $3;", u.Lang, u.LastMessageID, u.ID)
	if err != nil {
		return err
	}
	return nil
}
