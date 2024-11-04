package main

import (
	"log"
	"net/http"
	"posts/db"
	"posts/pages"
)

type Posts struct {
	Template string
	Id       uint
	Author   string
}

func InitPages() error {

	return nil
}

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	db.RunMigrations(database)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/login", pages.LoginPage)
	mux.HandleFunc("/", pages.MainPage)
	mux.HandleFunc("/post", pages.PostPage)

	http.ListenAndServe(":8010", mux)
}