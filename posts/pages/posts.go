package pages

import (
	"html/template"
	"log"
	"net/http"
	"posts/db"
	"strconv"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getid, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.NotFound(w, r)
			return
		}

		var post db.Posts
		result := db.DB.Where("ID = ?", getid).First(&post)
		if result.Error != nil {
			http.Error(w, "Такого поста не существует", http.StatusNotFound)
			return
		}
		tmpl, err := template.ParseFiles("./ui/static/html/posts.html", "./ui/static/html/base.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, post)
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}

}
