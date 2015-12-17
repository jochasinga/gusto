package routers

import (
        "github.com/gorilla/mux"

        "github.com/jochasinga/cli/secondapp/controllers"
)

var cbroker = controllers.NewBroker()

func (d *Delegate) HandleRoutes() *mux.Router {

        router := mux.NewRouter()
        router.HandleFunc("/", cbroker.HelloHandler)
        router.HandleFunc("/users", cbroker.GetUsersHandler)
        router.HandleFunc("/users/create", cbroker.CreateUserHandler)

        return router
}
