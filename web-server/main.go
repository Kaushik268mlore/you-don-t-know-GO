package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseeForm() err %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfful")
	name := r.FormValue("name")
	address := r.FormValue("Address")
	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "Address =%s\n", address)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not suppoted", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello Peeps!!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting the server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
