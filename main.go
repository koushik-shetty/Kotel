package main

import (
	a "TWLibrary/app"
	h "TWLibrary/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("****    ThoughtWorks Library application developed by Koushik Shetty[17801]  *****\n\n")

	var config a.AppConfig
	config.InitApp()

	h.Handlers("/", h.DefaultHandler)

	http.ListenAndServe(":5004", nil)
	return
}
