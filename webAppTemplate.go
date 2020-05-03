package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type NewsGroup struct{
	Title string
	News string
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/templ", templateHandler)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Welcome</h1>")
	fmt.Fprintf(w, "Hi, This is our first Handler!!")
}

func templateHandler(w http.ResponseWriter, r *http.Request){
	p := NewsGroup {Title: "World COVID-19 Pandemic", News: "some news"}
	t, _ := template.ParseFiles("MyTemplate.html")
	fmt.Println(t.Execute(w, p))
}