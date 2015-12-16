package models

type Delegate struct {}

// Register model functions or ORM here
type Broker interface {
        CreateUser(string, string, int)
        GetAllUsers() []User
        NewUser() *User
}

func NewBroker() Broker {
        return Broker(new(Delegate))
}
