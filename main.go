package main

import (
	"net/http"

	"github.com/sasasaiki/gokigen/src/go/router"
)

func main() {
	r := router.CreateRoute(router.NewProdRoutingHandlers())
	http.ListenAndServe(":8080", r)
}
