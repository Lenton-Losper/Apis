package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title "`
	Desc    string `json:"desc"`
	Content string `json:"*content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Practicing Api's", Desc: "Practicing Api's using Go", Content: "This is COntent fisrt of many"},
	}

	fmt.Println("Articles  Hit")
	fmt.Fprintf(w, "Aticles Page")
	json.NewEncoder(w).Encode(articles)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing if homepage is live")
}

func handleReq() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	handleReq()
}
