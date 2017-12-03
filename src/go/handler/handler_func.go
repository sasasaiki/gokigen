package gokigen

import (
	"net/http"
)

//HandlerFuncI ハンドリングすべき全てのfuncを持つ
type HandlerFuncI interface {
	add(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
	get(w http.ResponseWriter, r *http.Request)
}

//MyHandlerFunc ハンドリングするfuncとその情報を持つ
type MyHandlerFunc struct {
	f         func(w http.ResponseWriter, r *http.Request)
	path      string
	methods   []string
	needLogin bool
}

// 複数のエンドポイントで共有させたいオブジェクトとかもたせる
type ProdHandlerFunc struct {
}
