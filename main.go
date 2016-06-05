package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jochasinga/gusto/brokers/mbroker"
	"github.com/jochasinga/gusto/brokers/rbroker"
)

var (
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

	del := R.CarryVal()
	fmt.Println(del.M)
	router := R.HandleRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
