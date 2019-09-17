package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func Init() {
	fmt.Println("initialized mux router")
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	Router = router

}
