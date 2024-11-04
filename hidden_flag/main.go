package main

import (
	"html/template"
	"log"
	"net/http"
	"slices"
	"strconv"
)

type Data struct {
	TmpId int
}

var id []int

func InitId() {
	for i := 1; i < 1000; i++ {
		id = append(id, i)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	getid, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || !slices.Contains(id, getid) {
		http.NotFound(w, r)
		return
	}

	data := Data{
		TmpId: getid,
	}

	tmpl, err := template.ParseFiles("./ui/html/base.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func main() {
	InitId()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.ListenAndServe(":8020", mux)
}
