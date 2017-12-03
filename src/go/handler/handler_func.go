package handler

import (
	"net/http"
)

//====新しいハンドラーFuncは以下に追加====

//HandlerFuncI ハンドリングすべき全てのfuncを持つ。ハンドリングするfuncを増やす場合は追加
type HandlerFuncI interface {
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// ProdHandlerFunc 本番用。複数のエンドポイントで共有させたいオブジェクトとかもたせる。DBのコネクションとか？
type ProdHandlerFunc struct {
}

// 開発用などあれば以下に追加
