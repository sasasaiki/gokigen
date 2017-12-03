package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

//====新しいハンドラーは以下に追加===

//MyhandlerList 全てのHandlerを持つ。ハンドラーを増やす場合は追加
type MyhandlerList struct {
	Index http.Handler
}

//NewProdMyHandlerList prod用のHandlerリストを作る
func NewProdMyHandlerList() *MyhandlerList {
	return &MyhandlerList{
		Index: &templeteHandler{FileName: "main/index.html"},
	}
}

//========================================================

func (t *templeteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("indexにアクセス")
	//一度だけテンプレートを読み込む
	t.Once.Do(func() {
		t.Templ =
			template.Must(template.ParseFiles(filepath.Join("views",
				t.FileName)))
	})

	e := t.Templ.Execute(w, nil)

	if e != nil {
		fmt.Println("テンプレートの読み込みに失敗しています")
	}
}

func (lh logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	lh.nextHandler.ServeHTTP(w, r)
}

func (nh needLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("ここでログインのチェックを行いログインしていなかったらリダイレクトするような処理を書きます")
	nh.nextHandler.ServeHTTP(w, r)
}

//NewLogHandler 処理の前にログを吐くようにする
func NewLogHandler(h http.Handler) http.Handler {
	lh := logHandler{&decoratorHandler{nextHandler: h}}
	return &lh
}

//NewAuthHandler 処理の前にログインしているかチェックする
func NewAuthHandler(h http.Handler) http.Handler {
	nh := needLoginHandler{&decoratorHandler{nextHandler: h}}
	return &nh
}

type decoratorHandler struct {
	nextHandler http.Handler
}
type logHandler struct {
	*decoratorHandler
}
type needLoginHandler struct {
	*decoratorHandler
}

//templeteHandler htmlTempleteをを一度だけ読み込むハンドラー
type templeteHandler struct {
	Once     sync.Once
	FileName string
	Templ    *template.Template
}
