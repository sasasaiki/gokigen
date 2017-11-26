package imageFileServer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

//Myhandler ハンドリングすべき全てのfuncを持つ
type Myhandler struct {
	h       http.Handler
	path    string
	methods []string
}

//NewLogHandler 処理の前にログを吐くようにする
func NewLogHandler(h http.Handler) http.Handler {
	lh := logHandler{&decoratorHandler{nextHandler: h}}
	return &lh
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

func (lh *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	lh.nextHandler.ServeHTTP(w, r)
}

type decoratorHandler struct {
	nextHandler http.Handler
}
type logHandler struct {
	*decoratorHandler
}

type templeteHandler struct {
	once     sync.Once
	fileName string
	templ    *template.Template
}
