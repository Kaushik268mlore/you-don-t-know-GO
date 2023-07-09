package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.db

type user struct {
	gorm.model
	Name  string
	Email string
}

func allusers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "All users endpoint HIT")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the DB")
	}
	defer db.close()
	var users []user
	db.find(&users)
	json.NewEncoder(w).Encode(users)
}
func newuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user endpoint HIT")
}
func deleteuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete user endpoint HIT")
}
func updateuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user endpoint HIT")
}
func initmigrate() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the DB")
	}
	defer db.close()
	db.AutoMigrate(&user{})
}
