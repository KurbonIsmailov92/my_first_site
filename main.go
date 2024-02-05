package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static/css/main.css"))))
	http.HandleFunc("/", Index)
	http.HandleFunc("/create/", Create)
	http.HandleFunc("/save_article/", SaveArticle)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("tmp/index.html",
		"tmp/header.html",
		"tmp/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	template.ExecuteTemplate(w, "index", nil)

}

func Create(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("tmp/create.html",
		"tmp/header.html",
		"tmp/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	template.ExecuteTemplate(w, "create", nil)

}

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")

}

func main() {
	HandleFunc()
}
