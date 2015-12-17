package controllers

import (
        "encoding/json"
        "fmt"
        "net/http"

        "github.com/jochasinga/cli/secondapp/models"
)

var mbroker = models.NewBroker()

// Make all handlers methods of Delegate
func (d *Delegate) HelloHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello Gust!")
}

func (d *Delegate) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
        users := mbroker.GetAllUsers()
        uj, _ := json.Marshal(users)
        fmt.Fprintln(w, string(uj))
}

func (d *Delegate) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
        user := mbroker.NewUser()
        json.NewDecoder(r.Body).Decode(user)
        fmt.Println("CreateUserHandler here")
        mbroker.CreateUser(user.Username, user.Email, user.Age)
}
