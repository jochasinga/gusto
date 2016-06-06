package rbroker

import (
	"github.com/gorilla/mux"
	"github.com/jochasinga/gusto/routers"
)

type RBroker interface {
	Add(string, string)
	Del(string)
	Get(string) string
	Set(string, string)
	Vars() *routers.VarSet

	HandleRoutes() *mux.Router
}

func New() RBroker {
	delegate := &routers.Delegate{
		G: make(map[string][]string),
		VarSet: &routers.VarSet{
			Num:  routers.Num,
			Msg:  routers.Msg,
			Nums: routers.Nums,
		},
	}
	return RBroker(delegate)
}
