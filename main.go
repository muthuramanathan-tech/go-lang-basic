package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	//"github.com/gorilla/mux"
)

type Article struct{
	Title string `json:title"`
	Desc string `json:desc"`
	Content string `json:content`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title: "Test title", Desc: "test desc", Content:"test content"},
	}
	fmt.Println("EndPoint hit : All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint")
}

func handleRequest(){

	//incoming request and routing to approprate handler. For mor info https://www.gorillatoolkit.org/pkg/mux
	// myRouter := mux.NewRouter().StrictSlash(true)

	//Use Http endpoint
	 http.HandleFunc("/", homePage)
	 http.HandleFunc("/articles", allArticles)
	 log.Fatal(http.ListenAndServe(":8081", nil))

	//Use Mux Router endpoint
	/*myRouter.HandleFunc("/articles", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter)) */
	
}

func main(){
	handleRequest()
}