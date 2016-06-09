package routers

import (
	"github.com/gorilla/mux"

	"github.com/jochasinga/gusto/brokers/cbroker"
)

var C = cbroker.New()

var (
	Num  int    = 3
	Msg  string = "Hello"
	Nums []int  = []int{1, 2}
)

func (d *Delegate) HandleRoutes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", C.HelloHandler)
	router.HandleFunc("/users", C.GetUsersHandler)
	router.HandleFunc("/users/create", C.CreateUserHandler)

	cv := C.Vars()
	cv.MyFunc()
	fmt.Println(cv.MyParam)
	for _, v := range cv.MySlice {
		fmt.Println(v)
	}
	
	return router
}
