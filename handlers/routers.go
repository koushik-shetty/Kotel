package handlers

import (
	v "TWLibrary/views"
	"net/http"
)

type HanderFunction func(w http.ResponseWriter, r *http.Request)

func Handlers(patt string, handler HanderFunction) {
	http.HandleFunc(patt, handler)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(v.Index))
}
