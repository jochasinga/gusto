package models

type User struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Age      int    `json:"age"`
}

var Users []User

func (d *D) CreateUser(username, email string, age int) {
        Users = append(Users, User{
                Username: username,
                Email: email,
                Age: age,
        })
}

func (d *D) GetAllUsers() []User {
        return Users
}

func (d *D) NewUser() *User {
        return new(User)
}
