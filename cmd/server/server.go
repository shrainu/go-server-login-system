package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func NewUser(username, password string) User {

	return User{
		Id:       uuid.New(),
		Username: username,
		Password: password,
	}
}

type UserDB struct {
	Users []User `json:"users"`
}

func (db *UserDB) GetUserMap() map[uuid.UUID]User {

	users := map[uuid.UUID]User{}
	for _, v := range db.Users {
		users[v.Id] = v
	}

	return users
}

func GetDB() *UserDB {

	data, err := ioutil.ReadFile("database/users.json")
	if err != nil {
		log.Println(err)
		return nil
	}

	var db UserDB
	err = json.Unmarshal(data, &db)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &db
}

func ServeFile(rw http.ResponseWriter, r *http.Request) {

	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	url := path.Join(cwd, r.URL.EscapedPath())
	if url == cwd {
		url += "/template/home.html"
	}

	fmt.Println("Serving file:", url)

	http.ServeFile(rw, r, url)
}

func ServeHome(rw http.ResponseWriter, r *http.Request) {

	if r.URL.String() == "/home" {
		r.URL.Path = "/template/home.html"
		ServeFile(rw, r)
	}

	switch r.Method {
	case "GET":

		if r.URL.Query().Get("command") == "get-all-users" {

			rw.Header().Set("Content-Type", "application/json")

			userDB := GetDB()
			if userDB == nil {
				return
			}

			err := json.NewEncoder(rw).Encode(userDB)
			if err != nil {
				log.Println(err)
			}
		}
	default:
		log.Println("Unhandled request, method:", r.Method)
	}
}
