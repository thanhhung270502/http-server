package main

import (
	"net/http"

	"github.com/gowebexample/http-server/api"
)
// abc
func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}