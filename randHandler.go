package main

import (
	"html/template"
	//"math/rand"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("randTmpl.html")

	daysOfWeek := []string{"Mon.", "Tsu.", "Wed.", "Thu.", "Fri.", "Sat.", "Sun."}
	t.Execute(w, daysOfWeek)
	//t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
