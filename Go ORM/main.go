package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	rter := mux.NewRouter().StrictSlash(true)
	rter.HandleFunc("/", helloworld).Methods("GET")
	rter.HandleFunc("/users", allusers).Methods("GET")
	rter.HandleFunc("/user/{name}/{email}", newuser).Methods("POST")
	rter.HandleFunc("/user/{name}", deleteuser).Methods("DELETE")
	rter.HandleFunc("/user/{name}/{email}", updateuser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", rter))
}
func main() {
	fmt.Println("main func")
	initmigrate()
	handleRequest()
}
func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
