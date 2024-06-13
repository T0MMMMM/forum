package forum

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)


func (E *Engine) DataBaseCreation() {
	sqlFile, err := ioutil.ReadFile("./serv/sql/forum.sql")
    if err != nil {fmt.Printf("Error reading SQL file: %v", err)}

    if err != nil {fmt.Printf("Error connecting to the database: %v", err)}
    
	commands := splitSQLCommands(string(sqlFile))

    for _, cmd := range commands {
        _, err := E.DataBase.Exec(cmd)
        if err != nil {
            log.Fatalf("Erreur d'exécution de la commande SQL: %v", err)
        }
    }

	E.DataBase.Close()

}

func splitSQLCommands(file string) []string {
	list := strings.Split(file, ";")
	list = list[:len(list)-1] 
	for i := 0; i < len(list); i++ {
		list[i] += ";"
		fmt.Println(list[i])
	}
	
	return list
}


