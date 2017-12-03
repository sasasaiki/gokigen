package main

import (
	"net/http"

	"github.com/sasasaiki/gokigen/src/go/handler"
)

func main() {
	r := gokigen.CreateRoute(gokigen.NewProdRoutingHandlers())
	http.ListenAndServe(":8080", r)
}
