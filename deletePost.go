package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}
