package routers

import (
	"github.com/gorilla/mux"

	"github.com/jochasinga/gusto/brokers/cbroker"
)

var C = cbroker.New()

func (d *Delegate) CarryVal() *Delegate {
	d.M = make(map[string]string)
	d.M["key"] = "value"
	return d
}

func (d *Delegate) HandleRoutes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", C.HelloHandler)
	router.HandleFunc("/users", C.GetUsersHandler)
	router.HandleFunc("/users/create", C.CreateUserHandler)

	return router
}
