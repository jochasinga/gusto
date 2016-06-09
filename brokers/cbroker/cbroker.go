package cbroker

import (
	"net/http"

	"github.com/jochasinga/gusto/controllers"
	"github.com/jochasinga/gusto/routers"
)

type CBroker interface {
	Add(string, string)
	Del(string)
	Get(string) string
	Set(string, string)
	Vars() *routers.VarSet

	HelloHandler(http.ResponseWriter, *http.Request)
	GetUsersHandler(http.ResponseWriter, *http.Request)
	CreateUserHandler(http.ResponseWriter, *http.Request)
}

func New() CBroker {
	delegate := &controllers.Delegate{
		G: make(map[string][]string),
		VarSet: &controllers.VarSet{
			MyFunc:  controllers.MyFunc,
			MyParam: controllerss.MyParam,
			MySlice: controllers.MySlice,
		},
	}
	return RBroker(delegate)
}
