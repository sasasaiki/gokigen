package gokigen

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *ProdHandlerFunc) add(w http.ResponseWriter, r *http.Request) {
}

func (h *ProdHandlerFunc) update(w http.ResponseWriter, r *http.Request) {

}

func (h *ProdHandlerFunc) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ProdHandlerFunc) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["name"])
}

func outputError(w *http.ResponseWriter, e error, message string) {
	io.WriteString(*w, e.Error())
	log.Println(message, " エラーが発生しました:", e)
}
