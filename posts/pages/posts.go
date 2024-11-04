package pages

import (
	"net/http"
	"strconv"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	getid, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	_ = getid

}
