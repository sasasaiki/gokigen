package gokigen

import (
	"net/http"
)

//MyHandlerFunc ハンドリングするfuncとその情報を持つ
type MyHandlerFunc struct {
	f         func(w http.ResponseWriter, r *http.Request)
	path      string
	methods   []string
	needLogin bool
}

// ProdHandlerFunc 本番用。複数のエンドポイントで共有させたいオブジェクトとかもたせる。DBのコネクションとか？
type ProdHandlerFunc struct {
}

// 開発用などあれば以下に追加
