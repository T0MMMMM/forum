package forum

import (
	"html/template"
	"net/http"
)


func (E *Engine) index(w http.ResponseWriter, r *http.Request) {

	E.templateRender(w, "index")

}

func (E *Engine) templateRender(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./serv/html/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, E)
}