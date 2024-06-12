package forum

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	_ "github.com/mattn/go-sqlite3"
)


func (E *Engine) DataBase() {
	sqlFile, err := ioutil.ReadFile("./serv/sql/forum.sql")
    if err != nil {fmt.Printf("Error reading SQL file: %v", err)}
	db, err := sql.Open("sqlite3", "./serv/data/data.db")
    if err != nil {fmt.Printf("Error connecting to the database: %v", err)}
    defer db.Close()
	db.Exec(string(sqlFile))

}
