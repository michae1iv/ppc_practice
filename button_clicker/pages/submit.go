package pages

import (
	"clicker/jwt_handler"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func HandleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"success": true,
			"count":   0,
			"flag":    nil,
		}

		jwtString, err := r.Cookie("jwt")
		if err != nil {
			log.Println(err.Error())
			response["success"] = false
			json.NewEncoder(w).Encode(response)
			return
		}
		token, err := jwt_handler.CheckToken(jwtString.Value)
		if err != nil {
			log.Println(err.Error())
			response["success"] = false
			json.NewEncoder(w).Encode(response)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		clickCount := claims["clicks"].(float64) + 1
		tokenString, err := jwt_handler.CreateToken(claims["sub"].(string), clickCount)
		if err != nil {
			log.Println(err.Error())
			response["success"] = false
			json.NewEncoder(w).Encode(response)
			return
		}
		jwt_handler.SetToken(w, tokenString)

		response["count"] = clickCount
		if clickCount >= 1000 {
			response["flag"] = "flag{So_t1red_of_suchthing3}"
		}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}
}
