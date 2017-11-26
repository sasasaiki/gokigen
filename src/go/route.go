package gokigen

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewProdHandler 本番用ハンドラーを作成
func NewProdHandler() (*HandlerFuncI, *[]Myhandler) {
	var hf HandlerFuncI
	hf = new(prodHandlerFunc)
	hs := []Myhandler{
		{
			h:       &templeteHandler{fileName: "main/index.html"},
			path:    "/index",
			methods: []string{"GET"},
		},
	}
	return &hf, &hs
}

//CreateRoute 渡されたhandlerとfuncについてrouteを設定する
func CreateRoute(hf *HandlerFuncI, hs *[]Myhandler) *mux.Router {
	r := mux.NewRouter()

	setAPIRoute(r, hf)
	setRouteExistHandler(r, hs)

	// 404のときのハンドラ
	//r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	return r
}

//Handlerが必要ないrouteの設定
func setAPIRoute(r *mux.Router, hp *HandlerFuncI) {
	h := *hp
	setHandler(r, "/save", h.add, "POST")
	setHandler(r, "/get/{name}/", h.get, "GET")
	setHandler(r, "/update", h.update, "PUT")
	setHandler(r, "/delete", h.delete, "DELETE")
}

func setHandler(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request), methods ...string) {
	lh := NewLogHandler(http.HandlerFunc(f))
	r.Handle(path, lh).Methods(methods...)
}

//Handlerが必要なrouteの設定
//templete読み込みなど
//TODO:structじゃなくて配列で持つようにしよう
func setRouteExistHandler(r *mux.Router, hs *[]Myhandler) {
	for _, h := range *hs {
		lh := NewLogHandler(h.h)
		r.Handle(h.path, lh).Methods(h.methods...)
	}
}
