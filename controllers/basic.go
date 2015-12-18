package controllers

import (
        "encoding/json"
        "fmt"
        "net/http"

        "github.com/jochasinga/cli/secondapp/brokers/mbroker"
)

var M = mbroker.New()

func (d *D) HelloHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello Gust!")
}

func (d *D) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
        users := M.GetAllUsers()
        uj, _ := json.Marshal(users)
        fmt.Fprintln(w, string(uj))
}

func (d *D) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
        user := M.NewUser()
        json.NewDecoder(r.Body).Decode(user)
        fmt.Println("CreateUserHandler here")
        M.CreateUser(user.Username, user.Email, user.Age)
}
