/*
Take the previous program and change it so that:
func main uses http.Handle instead of http.HandleFunc.
Contstraint: Do not change anything outside of func main.
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
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
