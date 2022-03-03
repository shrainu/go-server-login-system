package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

func checkValidUsername(username string) (bool, error) {

	if len(username) < 3 {
		return false, fmt.Errorf("username can't be shorter than 3 characters")
	}

	userDB := GetDB()

	for _, v := range userDB.Users {
		if v.Username == username {
			return false, fmt.Errorf("username already taken")
		}
	}

	valid := !strings.ContainsAny(
		username,
		"!$%^&*()@?><||{}';:\"[]#.,\\-+*/` ",
	)

	if !valid {
		return false, fmt.Errorf("username contains forbidden characters")
	}

	return true, nil
}

func checkValidPassword(password string) (bool, error) {

	if len(password) < 8 {
		return false, fmt.Errorf("password can't be shorter than 8 characters")
	}

	valid := !strings.ContainsAny(
		password,
		" ;:\"\\'",
	)

	if !valid {
		return false, fmt.Errorf("password can't contain ; : \" \\ '")
	}

	return true, nil
}

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

func AddUser(username, password string) User {

	user := NewUser(username, password)

	userDB := GetDB()
	if userDB == nil {
		return User{}
	}
	userDB.Users = append(userDB.Users, user)

	file, err := os.OpenFile("database/users.json", os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(userDB, "", "  ")
	if err != nil {
		log.Println(err)
	}

	file.Write(data)

	return user
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

	if r.URL.String() == "/home" && r.Method == "GET" {
		r.URL.Path = "/template/home.html"
		ServeFile(rw, r)
		return
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
	case "POST":

		rw.Header().Set("Content-Type", "application/json")

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
		}

		validUsername, err := checkValidUsername(user.Username)
		if !validUsername {

			rw.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintln(rw, err.Error())
			return
		}

		validPassword, err := checkValidPassword(user.Password)
		if !validPassword {

			rw.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintln(rw, err.Error())
			return
		}

		added := AddUser(user.Username, user.Password)

		// Users data shouldn't be returned in a normal application,
		// but its okay for ours since, this is done for educational purposes.
		_, err = fmt.Fprintf(
			rw,
			"{\"id\":\"%s\", \"username\":\"%s\", \"password\":\"%s\"}\n",
			added.Id.String(),
			added.Username,
			added.Password,
		)
		if err != nil {
			log.Println(err)
		}
	default:
		log.Println("Unhandled request, method:", r.Method)
	}
}
