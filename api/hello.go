package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
