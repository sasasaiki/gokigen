package main

import (
	"net/http"

	"github.com/sasasaiki/gokigen/src/go"
)

func main() {
	r := gokigen.CreateRoute(gokigen.NewProdHandler())
	//cssやjsを読み込めるようにするHandler
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	http.ListenAndServe(":8080", r)
}
