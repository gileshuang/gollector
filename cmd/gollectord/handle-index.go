package main

import (
	"html/template"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.URL.Path != "/" && r.URL.Path != "/index/" {
		http.NotFound(w, r)
		return
	}

	t, _ := template.ParseFiles("template/index.gtpl")
	t.Execute(w, r.URL.Path)
	return
}
