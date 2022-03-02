package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	username string
	password string
}

func serveFile(rw http.ResponseWriter, r *http.Request) {

	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	url := path.Join(cwd, r.URL.EscapedPath())

	http.ServeFile(rw, r, url)
}

func main() {

	port := ":8080"

	http.HandleFunc("/", serveFile)

	fmt.Printf("Server starting at port %v\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
