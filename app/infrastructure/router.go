package infrastructure

import (
	"../interfaces/controller"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func Init() {
	router := mux.NewRouter()

	conn := Connect()
	memoController := controller.NewMemoController(conn)
	router.HandleFunc("/memos", memoController.FindAll).Methods("GET")
	router.HandleFunc("/memos", memoController.Create).Methods("POST")
	Router = router

}
