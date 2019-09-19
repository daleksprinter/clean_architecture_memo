package main

import (
	"net/http"

	"github.com/daleksprinter/clean_todo/server/infrastructure"
)

func main() {
	infrastructure.Init()
	http.ListenAndServe(":8080", infrastructure.Router)
}
