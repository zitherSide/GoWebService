package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Fprintln(w, "form", r.Form)
	//fmt.Fprintln(w, "post from", r.PostForm)

	// file, _, err := r.FormFile("uploaded")
	// if err == nil {
	// 	data, err := ioutil.ReadAll(file)
	// 	if err == nil {
	// 		fmt.Fprintln(w, string(data))
	// 	}
	// }

	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintf(w, string(data))
		}
	}

	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, r.PostFormValue("hello"))
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>
	`

	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such Service. Finde another service.")
}

func headerExmaple(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Takuya Kotoyori",
		Threads: []string{"1st, 2nd, 3rd"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExmaple)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
