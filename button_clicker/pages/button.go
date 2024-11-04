package pages

import (
	"clicker/jwt_handler"
	"html/template"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func ButtonPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		jwtString, err := r.Cookie("jwt")
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		token, err := jwt_handler.CheckToken(jwtString.Value)
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		claims := token.Claims.(jwt.MapClaims)

		data := jwt_handler.Data{
			Clicks: claims["clicks"].(float64),
		}

		tmpl, err := template.ParseFiles("./ui/static/html/button.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}

}
