package main

import (
	"fs/ascii"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	Name string
}

func main() {
	http.HandleFunc("/", Just_display)
	http.HandleFunc("/result", Accept_the_file)
	http.ListenAndServe("localhost:8128", nil)
}

func Just_display(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	h, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	h.Execute(w, nil)
}

func Accept_the_file(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Header().Set("Allow", http.MethodPost)
		w.Write([]byte("Method not allowed!"))
		return
	}
	h, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	t, err0 := r.Form["inserttext"]
	if !err0 {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	str := strings.Join(t, " ")
	st := r.FormValue("textstyle")
	s, err1 := ascii.Ascii_art(str, st)
	if err1 == 500 {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}
	if err1 == 400 {
		w.WriteHeader(400)
		w.Header().Set("Allow", http.MethodPost)
		w.Write([]byte("400 - Bad Request"))
		return
	}
	res := Data{
		Name: s,
	}
	h.Execute(w, res)
}
