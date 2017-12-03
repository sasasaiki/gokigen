package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *ProdHandlerFunc) Add(w http.ResponseWriter, r *http.Request) {
}

func (h *ProdHandlerFunc) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ProdHandlerFunc) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ProdHandlerFunc) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["name"])
}

func outputError(w *http.ResponseWriter, e error, message string) {
	io.WriteString(*w, e.Error())
	log.Println(message, " エラーが発生しました:", e)
}
