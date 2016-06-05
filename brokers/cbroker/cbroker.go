package cbroker

import (
	"net/http"

	"github.com/jochasinga/gusto/controllers"
)

type CBroker interface {
	HelloHandler(http.ResponseWriter, *http.Request)
	GetUsersHandler(http.ResponseWriter, *http.Request)
	CreateUserHandler(http.ResponseWriter, *http.Request)
}

func New() CBroker {
	return CBroker(new(controllers.Delegate))
}
