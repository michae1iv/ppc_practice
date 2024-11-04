package pages

import (
	"html/template"
	"log"
	"net/http"
	"posts/db"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var posts []db.Posts
		if err := db.DB.Find(&posts).Error; err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl, err := template.ParseFiles("./ui/static/html/index.html", "./ui/static/html/base.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, posts)
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}
}
