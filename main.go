package main

import (
	"fmt"
	"log"
	"net/http"

	"login-system/cmd/server"
)

func main() {

	port := ":8080"

	http.HandleFunc("/home", server.ServeHome)
	http.HandleFunc("/", server.ServeFile)

	fmt.Printf("Server starting at port %v\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
