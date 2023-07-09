package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"Content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test Title ",
			Desc:    "test desc",
			Content: "Test content",
		},
	}
	fmt.Println("ENDPOINT HIT: ALL articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func testPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test POST ENDPOINT WORKED")
}
func HandleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homepage)
	myrouter.HandleFunc("/articles", allArticles).Methods("GET")
	myrouter.HandleFunc("/articles", testPOST).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myrouter))
}
func main() {
	HandleRequests()
}
