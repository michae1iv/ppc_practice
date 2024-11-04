package main

import (
	"clicker/pages"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/submit", pages.HandleSubmit)
	mux.HandleFunc("/", pages.RegistrationPage)
	mux.HandleFunc("/button", pages.ButtonPage)

	http.ListenAndServe(":8080", mux)
}
