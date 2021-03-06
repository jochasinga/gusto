package mbroker

import (
	"github.com/jochasinga/gusto/models"
)

type MBroker interface {
	CreateUser(string, string, int)
	GetAllUsers() []models.User
	NewUser() *models.User
}

func New() MBroker {
	return MBroker(new(models.Delegate))
}
