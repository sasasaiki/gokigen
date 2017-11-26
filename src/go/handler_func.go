package webServer

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//HandlerFuncI ハンドリングすべき全てのfuncを持つ
type HandlerFuncI interface {
	add(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
	get(w http.ResponseWriter, r *http.Request)
}

func (h *prodHandlerFunc) add(w http.ResponseWriter, r *http.Request) {
}

func (h *prodHandlerFunc) update(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlerFunc) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlerFunc) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["name"])
}

func outputError(w *http.ResponseWriter, e error, message string) {
	io.WriteString(*w, e.Error())
	log.Println(message, " エラーが発生しました:", e)
}

type prodHandlerFunc struct {
}
