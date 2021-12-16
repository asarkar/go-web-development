/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func root(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, "root")
}

func dog(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, "dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) == 3 {
		writeResponse(w, paths[2])
	}
}

func writeResponse(w http.ResponseWriter, s string) {
	_, err := io.WriteString(w, s+"\n")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
