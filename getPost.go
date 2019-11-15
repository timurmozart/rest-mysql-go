package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(post)
}
