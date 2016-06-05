package rbroker

import (
	"github.com/gorilla/mux"
	"github.com/jochasinga/gusto/routers"
)

type RBroker interface {
	HandleRoutes() *mux.Router
	CarryVal() *routers.Delegate
}

func New() RBroker {
	return RBroker(new(routers.Delegate))
}
