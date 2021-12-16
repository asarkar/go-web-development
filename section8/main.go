/* Using cookies, track how many times a user has been to your website domain. */
package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handle)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("visit")
	if err != nil {
		c = &http.Cookie{Name: "visit", Value: "0"}
	}
	visit, _ := strconv.Atoi(c.Value)
	c.Value = strconv.Itoa(visit + 1)
	http.SetCookie(w, c)
	_, err = io.WriteString(w, c.Value)
	if err != nil {
		log.Fatal(err)
	}
}
