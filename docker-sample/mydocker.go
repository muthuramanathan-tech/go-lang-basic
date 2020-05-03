package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("Docker GoLang")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello Worldas")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}