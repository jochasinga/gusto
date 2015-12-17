package main

import (
        "log"
        "net/http"

        "github.com/jochasinga/cli/secondapp/controllers"
        "github.com/jochasinga/cli/secondapp/models"
        "github.com/jochasinga/cli/secondapp/routers"
)

var (
        cbroker = controllers.NewBroker()
        mbroker = models.NewBroker()
        rbroker = routers.NewBroker()
)

// REMOVE: Used in populating sample data only
func init() {
        mbroker.CreateUser("kgibran", "kahlil_g@mail.com", 48)
        mbroker.CreateUser("robpike", "rob_pike@gmail.com", 59)
        mbroker.CreateUser("kenthompson", "kthompson@gmail.com", 72)
        mbroker.CreateUser("pchasinga", "jo.chasinga@gmail.com", 32)
}

func main() {

        router := rbroker.HandleRoutes()
        log.Fatal(http.ListenAndServe(":8080", router))
}
