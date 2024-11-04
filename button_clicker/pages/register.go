package pages

import (
	"clicker/jwt_handler"
	"html/template"
	"log"
	"net/http"
)

func RegistrationPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/static/html/register.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		tokenString, err := jwt_handler.CreateToken(r.FormValue("username"), 0)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		jwt_handler.SetToken(w, tokenString)
		http.Redirect(w, r, "/button", http.StatusSeeOther)
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}
}
