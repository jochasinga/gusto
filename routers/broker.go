package routers

import (
        "github.com/gorilla/mux"
)

type Delegate struct {}

// Register model functions or ORM here
type Broker interface {
        HandleRoutes() *mux.Router
}

func NewBroker() Broker {
        return Broker(new(Delegate))
}
