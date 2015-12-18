package rbroker

import (
        "github.com/gorilla/mux"
        "github.com/jochasinga/cli/secondapp/routers"
)

type RBroker interface {
        HandleRoutes() *mux.Router
}

func New() RBroker {
        return RBroker(new(routers.D))
}
