package gokigen

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewProdRoutingHandlers 本番用ハンドラーを作成
func NewProdRoutingHandlers() (*[]MyHandlerFunc, *[]MyHandler) {
	hf := NewHandlerFuncs(new(ProdHandlerFunc))
	hs := NewHandlers(NewProdMyHandlerList())
	return &hf, &hs
}

//CreateRoute 渡されたhandlerとfuncについてrouteを設定する
func CreateRoute(hf *[]MyHandlerFunc, hs *[]MyHandler) *mux.Router {
	r := mux.NewRouter()

	//cssやjsを読み込めるようにするHandler
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	setFuncsRoute(r, hf)
	setHandlersRoute(r, hs)
	// 404のときのハンドラ
	//r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	return r
}

//Handlerが必要ないrouteの設定
func setFuncsRoute(r *mux.Router, hf *[]MyHandlerFunc) {
	for _, h := range *hf {
		setRoute(r, h.path, http.HandlerFunc(h.f), h.needLogin, h.methods...)
	}
}

//Handlerが必要なrouteの設定
//templete読み込みなど
func setHandlersRoute(r *mux.Router, hs *[]MyHandler) {
	for _, h := range *hs {
		setRoute(r, h.path, h.h, h.needLogin, h.methods...)
	}
}

//新しくHandlerをデコレーションする必要がある場合はここでやる
func setRoute(r *mux.Router, path string, h http.Handler, needLogin bool, methods ...string) {
	result := h
	if needLogin {
		result = NewAuthHandler(result)
	}
	result = NewLogHandler(result)
	r.Handle(path, result).Methods(methods...)
}
