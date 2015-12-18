package main

import (
        "log"
        "net/http"

        //"github.com/jochasinga/cli/secondapp/controllers"
        //"github.com/jochasinga/cli/secondapp/models"
        //"github.com/jochasinga/cli/secondapp/routers"
        "github.com/jochasinga/cli/secondapp/cbroker"
        "github.com/jochasinga/cli/secondapp/mbroker"
        "github.com/jochasinga/cli/secondapp/rbroker"
)

var (
        C = cbroker.New()
        M = mbroker.New()
        R = rbroker.New()
)

// REMOVE: Used in populating sample data only
func init() {
        M.CreateUser("kgibran", "kahlil_g@mail.com", 48)
        M.CreateUser("robpike", "rob_pike@gmail.com", 59)
        M.CreateUser("kenthompson", "kthompson@gmail.com", 72)
        M.CreateUser("pchasinga", "jo.chasinga@gmail.com", 32)
}

func main() {

        router := R.HandleRoutes()
        log.Fatal(http.ListenAndServe(":8080", router))
}
