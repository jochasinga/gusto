package routers

import (
        "github.com/gorilla/mux"

        "github.com/jochasinga/cli/secondapp/brokers/cbroker"
)

var C = cbroker.New()

func (d *D) HandleRoutes() *mux.Router {

        router := mux.NewRouter()
        router.HandleFunc("/", C.HelloHandler)
        router.HandleFunc("/users", C.GetUsersHandler)
        router.HandleFunc("/users/create", C.CreateUserHandler)

        return router
}
