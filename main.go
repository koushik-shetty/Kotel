package main

import (
	a "TWLibrary/app"
	h "TWLibrary/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("****    ThoughtWorks Library application developed by Koushik Shetty[17801]  *****\n")

	AppConfig := a.InitApp()

	fmt.Printf("App port : %s\n", AppConfig.GetConfig().Port)
	h.Handlers("/", h.DefaultHandler)
	h.Handlers("/public/", h.PublicHandler)

	http.ListenAndServe(":"+AppConfig.GetConfig().Port, nil)
	return
}
