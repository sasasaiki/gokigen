package router

import (
	"net/http"

	"github.com/sasasaiki/gokigen/src/go/handler"
)

//ルートの追加はこのファイルで行う

//NewHandlerFuncs funcの設定を配列としてもつ。新しくハンドリングするときはここに追加。
func NewHandlerFuncs(h handler.HandlerFuncI) []MyHandlerFunc {
	return []MyHandlerFunc{
		{
			f:         h.Add,
			path:      "/save",
			methods:   []string{"POST"},
			needLogin: true,
		},
		{
			f:         h.Get,
			path:      "/get/{firstName}/{lastName}",
			methods:   []string{"GET"},
			needLogin: false,
		},
		{
			f:         h.Update,
			path:      "/update",
			methods:   []string{"PUT"},
			needLogin: true,
		},
		{
			f:         h.Delete,
			path:      "/delete",
			methods:   []string{"DELETE"},
			needLogin: true,
		},
	}
}

//NewHandlers Handlerの設定を配列としてもつ。新しくハンドリングするときはここに追加。
func NewHandlers(hl *handler.MyhandlerList) []MyHandler {
	return []MyHandler{
		{
			h:         hl.Index,
			path:      "/",
			methods:   []string{"GET"},
			needLogin: false,
		},
	}
}

//MyHandlerFunc ハンドリングするfuncとその情報を持つ
type MyHandlerFunc struct {
	f         func(w http.ResponseWriter, r *http.Request)
	path      string
	methods   []string
	needLogin bool
}

//MyHandler Handlerとその設置
type MyHandler struct {
	h         http.Handler
	path      string
	methods   []string
	needLogin bool
}
