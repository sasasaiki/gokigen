package gokigen

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

//MyHandler Handlerとその設置
type MyHandler struct {
	h         http.Handler
	path      string
	methods   []string
	needLogin bool
}

//NewProdMyHandlerList prod用のHandlerリストを作る
func NewProdMyHandlerList() *MyhandlerList {
	return &MyhandlerList{
		index: &templeteHandler{fileName: "main/index.html"},
	}
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

func (t *templeteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("indexにアクセス")
	//一度だけテンプレートを読み込む
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("views",
				t.fileName)))
	})

	e := t.templ.Execute(w, nil)

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

type decoratorHandler struct {
	nextHandler http.Handler
}
type logHandler struct {
	*decoratorHandler
}
type needLoginHandler struct {
	*decoratorHandler
}

//====新しいハンドラーは以下に追加===

type templeteHandler struct {
	once     sync.Once
	fileName string
	templ    *template.Template
}
