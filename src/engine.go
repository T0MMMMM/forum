package forum

import (
	"fmt"
	// "math/rand"
	"net/http"
)

type Engine struct {
	Port		string
}

func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())

	E.Port = ":8080"

}


func (E *Engine) Run() {
	E.Init()

	fs := http.FileServer(http.Dir("serv"))
	http.Handle("/serv/", http.StripPrefix("/serv/", fs))

	http.HandleFunc("/", E.index)

	fmt.Println("(http://localhost:8080) - Serveur started on port", E.Port)

	http.ListenAndServe(E.Port, nil)
}