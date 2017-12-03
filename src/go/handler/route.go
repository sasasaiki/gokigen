package gokigen

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewHandlerFuncs funcの設定を配列としてもつ。新しくハンドリングするときはここに追加。
func NewHandlerFuncs(h HandlerFuncI) []MyHandlerFunc {
	return []MyHandlerFunc{
		{
			f:         h.add,
			path:      "/save",
			methods:   []string{"POST"},
			needLogin: true,
		},
		{
			f:         h.get,
			path:      "/get/{name}/",
			methods:   []string{"GET"},
			needLogin: false,
		},
		{
			f:         h.update,
			path:      "/update",
			methods:   []string{"PUT"},
			needLogin: true,
		},
		{
			f:         h.delete,
			path:      "/delete",
			methods:   []string{"DELETE"},
			needLogin: true,
		},
	}
}

//NewHandlers Handlerの設定を配列としてもつ。新しくハンドリングするときはここに追加。
func NewHandlers(hl *MyhandlerList) []MyHandler {
	return []MyHandler{
		{
			h:         hl.index,
			path:      "/index",
			methods:   []string{"GET"},
			needLogin: false,
		},
	}
}

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

func setRoute(r *mux.Router, path string, h http.Handler, needLogin bool, methods ...string) {
	result := h
	if needLogin {
		result = NewAuthHandler(result)
	}
	result = NewLogHandler(result)
	r.Handle(path, result).Methods(methods...)
}
