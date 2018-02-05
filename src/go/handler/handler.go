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

//Handlers 全てのHandlerを持つ。ハンドラーを増やす場合は追加
type Handlers struct {
	Index http.Handler
}

//NewHandlers Handlerの設定を配列としてもつ。新しくハンドリングするときはここに追加。
func NewHandlers(hl *Handlers) []Handler {
	return []Handler{
		{
			Handler: hl.Index,
			Conf: &HandlingConf{
				Path:      "/",
				Methods:   []string{"GET"},
				NeedLogin: false,
			},
		},
	}
}

//NewProdMyHandlerList prod用のHandlerリストを作る
func NewProdMyHandlerList() *Handlers {
	return &Handlers{
		Index: &templateHandler{Template: &Template{FileName: "main/index.html"}},
	}
}

//Handler Handlerとその設定
type Handler struct {
	Handler http.Handler
	Conf    *HandlingConf
}

//HandlingConf handlerの設定
type HandlingConf struct {
	Path      string
	Methods   []string
	NeedLogin bool
}

//========================================================

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	execTemplate(w, t.Template, nil)
}

func execTemplate(w http.ResponseWriter, t *Template, param interface{}) {
	log.Println(t.FileName + "にアクセス")
	t.Once.Do(func() {
		t.Templ =
			template.Must(template.ParseFiles(filepath.Join("views",
				t.FileName)))
	})
	e := t.Templ.Execute(w, param)

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

//Template templateHandlerに持たせる
type Template struct {
	Once     sync.Once
	FileName string
	Templ    *template.Template
}

//templateHandler htmlTemplateを一度だけ読み込むハンドラー
//goのテンプレートに値を渡したいときはこいつを埋め込んで値のstructも持たせ
//ServeHTTPを実装する
type templateHandler struct {
	*Template
}
