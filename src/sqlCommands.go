package forum

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
This function allows you to create all the bases of the program if they do not yet exist
*/
func (E *Engine) DataBaseCreation() {
	sqlFile, err := ioutil.ReadFile("./serv/sql/forum.sql")
	if err != nil {fmt.Printf("Error reading SQL file: %v", err)}
	if err != nil {fmt.Printf("Error connecting to the database: %v", err)}

	commands := splitSQLCommands(string(sqlFile))

	for _, cmd := range commands {
		_, err := E.DataBase.Exec(cmd)
		if err != nil {log.Fatalf("Erreur d'exécution de la commande SQL: %v", err)}
	}
	_, err = E.DataBase.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {log.Fatal(err)}

	// Catégories Creation
	// E.ExecuteSQL("INSERT INTO `categories` (name, description) VALUES ('Developpement', 'Developpement informatique')")
	// E.ExecuteSQL("INSERT INTO `categories` (name, description) VALUES ('Cinema', 'Tous les sujets possibles du cinéma !')")
	// E.ExecuteSQL("INSERT INTO `categories` (name, description) VALUES ('Jeux Vidéos', 'Parlons de gaming !')")
}


/* 
This function helps separate different SQL commands so that they can be executed correctly and line by line.
*/
func splitSQLCommands(file string) []string {
	list := strings.Split(file, ";")
	list = list[:len(list)-1]
	for i := 0; i < len(list); i++ {
		list[i] += ";"
	}
	return list
}

/*
This function allows you to execute a SQL line
*/
func (E *Engine) ExecuteSQL(cmd string) error {
	_, err := E.DataBase.Exec(cmd)
	if err != nil {
		log.Fatalf("Erreur d'exécution de la commande SQL: %v", err)
		return err
	}
	return nil
}

/*
This function allows you to execute a SQL line which allows you to retrieve data
*/
func (E *Engine) QuerySQL(cmd string) *sql.Rows {
	data, err := E.DataBase.Query(cmd)
	if err != nil {
		log.Fatalf("Erreur d'exécution de la commande SQL: %v", err)
	}
	return data
}

