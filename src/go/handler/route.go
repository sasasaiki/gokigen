package gokigen

import (
	"net/http"
)

//ルートの追加はこのファイルで行う

//MyhandlerList 全てのHandlerを持つ。ハンドラーを増やす場合は追加
type MyhandlerList struct {
	index http.Handler
}

//HandlerFuncI ハンドリングすべき全てのfuncを持つ。ハンドリングするfuncを増やす場合は追加
type HandlerFuncI interface {
	add(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
	get(w http.ResponseWriter, r *http.Request)
}

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
