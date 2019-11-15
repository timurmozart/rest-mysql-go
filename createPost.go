package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("insert into posts(title) values(?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]

	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New post created!")
}
