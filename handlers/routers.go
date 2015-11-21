package handlers

import (
	v "TWLibrary/views"
	"fmt"
	"html/template"
	"net/http"
)

type HanderFunction func(w http.ResponseWriter, r *http.Request)

func Handlers(patt string, handler HanderFunction) {
	http.HandleFunc(patt, handler)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.New("index.html").Parse(v.Index)

	options := &v.IndexOptions{
		Title: "TW Library",
	}
	e = t.Execute(w, options)
	if e != nil {
		fmt.Printf("Error Generating template : %v", e)
	}
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Contrl", "private, max-age=31536000")
	fmt.Printf("url paths : %v\n", r.URL.Path)
	http.ServeFile(w, r, r.URL.Path[1:])
}
