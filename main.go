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
	router := R.HandleRoutes()
	R.Add("name1", "Jordan")
	R.Add("name2", "Mike")
	fmt.Println(R.Get("name1"))

	routerVars := R.Vars()
	fmt.Println(routerVars.Num)
	fmt.Println(routerVars.Nums)
	log.Fatal(http.ListenAndServe(":8080", router))
}
