package main

import (
	"net/http"

	"./infrastructure"
)

func main() {
	infrastructure.Init()
	http.ListenAndServe(":8080", infrastructure.Router)
}
