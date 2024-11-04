package pages

import (
	"html/template"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/static/html/index.html", "./ui/static/html/base.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, nil)
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}
}
