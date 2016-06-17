package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	con "github.com/koushik-shetty/Kotel/constants"
	v "github.com/koushik-shetty/Kotel/views"
)

type HanderFunction func(w http.ResponseWriter, r *http.Request)

func Handlers(patt string, handler HanderFunction) {
	http.HandleFunc(patt, handler)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.New("index.html").Parse(v.Index)

	options := &v.IndexOptions{
		Title: con.AppName,
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("success:false"))
		fmt.Printf("err:%v", err)
	}
	userid := r.FormValue("userid")
	pass := r.FormValue("password")
	fmt.Printf("user=%s,pass=%s", userid, pass)
	w.Write([]byte("success:true"))
}
