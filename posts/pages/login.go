package pages

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"posts/db"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/static/html/login.html", "./ui/static/html/base.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user db.Users
		result := db.DB.Where(&db.Users{Username: username, Password: password}).First(&user)
		if result.Error != nil {
			http.Error(w, "Неверное имя пользователя или пароль", http.StatusUnauthorized)
			return
		}
		fmt.Fprint(w, "flag{y0u_Ar3_REEEally_ppC_master_Congrats}")
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}
}
