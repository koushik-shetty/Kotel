package main

import (
	"fmt"
	"log"

	a "github.com/koushik-shetty/Kotel/app"
	c "github.com/koushik-shetty/Kotel/config"
	h "github.com/koushik-shetty/Kotel/handlers"
	"net/http"
)

func main() {
	fmt.Println("***************************************************************************\n")
	fmt.Println("*       Kotel(Virtual Story Wall) developed by Koushik Shetty[17801]      *\n")
	fmt.Println("*				koushik@thoughtworks.com/kou4307@gmail.com                 *\n")
	fmt.Println("*					https://github.com/koushik-shetty                      *\n")
	fmt.Println("***************************************************************************\n")

	AppConfig, err := a.InitApp()
	if err != nil {
		log.Fatalf("Error reading appconfig: %s", err)
	}

	fmt.Printf("App port : %s\n", AppConfig.GetPort())
	h.Handlers("/public/", h.PublicHandler)
	h.Handlers("/login", h.LoginHandler)
	h.Handlers("/", h.DefaultHandler)

	//TODO : handle logger error

	dbconf := c.DefaultDBConfig()
	db, _ := a.NewDatabase(dbconf)
	defer db.Database().Close()

	err = db.Database().Ping()

	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	http.ListenAndServe(":"+AppConfig.GetPort(), nil)
	return
}
