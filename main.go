package main

import (
        "log"
        "net/http"

        "./controllers"
        "./models"

        "github.com/gorilla/mux"
)

var (
        cbroker = controllers.NewBroker()
        mbroker = models.NewBroker()
)

// REMOVE: Used in populating sample data only
func init() {
        mbroker.CreateUser("kgibran", "kahlil_g@mail.com", 48)
        mbroker.CreateUser("robpike", "rob_pike@gmail.com", 59)
        mbroker.CreateUser("kenthompson", "kthompson@gmail.com", 72)
        mbroker.CreateUser("pchasinga", "jo.chasinga@gmail.com", 32)
}

func main() {

        router := mux.NewRouter()
        router.HandleFunc("/", cbroker.HelloHandler)
        router.HandleFunc("/users", cbroker.GetUsersHandler)
        router.HandleFunc("/users/create", cbroker.CreateUserHandler)

        log.Fatal(http.ListenAndServe(":8080", router))
}
