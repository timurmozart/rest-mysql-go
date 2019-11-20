package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Post our structure of posts
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// PageVariables date time
type PageVariables struct {
	Date string
	Time string
}

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "root:ASdf456+@tcp(mariadb:3306)/qwer")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	router.HandleFunc("/home", HomePage)

	http.ListenAndServe(":8000", router)

}
