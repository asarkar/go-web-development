/* Create a server that returns the URL of the GET request */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	log.Println("Server is running...")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	log.Println("Received request")

	uri := parseuri(conn)

	respond(conn, uri)
}

func parseuri(conn net.Conn) string {
	s := bufio.NewScanner(conn)
	var uri string
	firstline := true
	for s.Scan() {
		line := s.Text()
		if firstline { // GET / HTTP/1.0
			firstline = false
			tokens := strings.Fields(line)
			if len(tokens) == 3 && strings.HasPrefix(tokens[0], "GET") {
				uri = tokens[1]
			} else {
				line = ""
			}
		}
		if line == "" {
			break
		}
	}
	return uri
}

func respond(conn net.Conn, body string) {
	if body != "" {
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/plain\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	} else {
		fmt.Fprint(conn, "HTTP/1.1 400 Bad Request\r\n")
	}
	log.Println("Sent response")
}
