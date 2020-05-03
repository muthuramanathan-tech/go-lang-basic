package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/second", IndexSecondHndlr)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Welcome</h1>")
	fmt.Fprintf(w, "Hi, This is our first Handler!!")
}

func IndexSecondHndlr(w http.ResponseWriter, r *http.Request){
	fmt.Println("hiiddiid")
	fmt.Fprintf(w, "Hi, This is our Second Handler!!")
}