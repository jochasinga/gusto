package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jochasinga/gusto/brokers/mbroker"
)

var (
	M = mbroker.New()
	
	MyFunc = func() {
		fmt.Println("Funcky!")
	}
	MyParam = "Param"
	MySlice = []string{"foo", "bar", "baz"}
)

func (d *Delegate) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Gust!")
}

func (d *Delegate) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := M.GetAllUsers()
	uj, _ := json.Marshal(users)
	fmt.Fprintln(w, string(uj))
}

func (d *Delegate) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := M.NewUser()
	json.NewDecoder(r.Body).Decode(user)
	fmt.Println("CreateUserHandler here")
	M.CreateUser(user.Username, user.Email, user.Age)
}
