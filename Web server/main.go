package main

import (
	"my-project/api"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":3000", srv)
}
