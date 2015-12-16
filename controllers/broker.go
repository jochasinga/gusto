package controllers

import (
        "net/http"
)

// Delegate can be left empty or with optional parameters
type Delegate struct {}

// This is where you register all handler functions
type Broker interface {
        HelloHandler(http.ResponseWriter, *http.Request)
        GetUsersHandler(http.ResponseWriter, *http.Request)
        CreateUserHandler(http.ResponseWriter, *http.Request)
}

func NewBroker() Broker {
        return Broker(new(Delegate))
}
