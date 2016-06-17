package main

import (
	"fmt"
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

	AppConfig := a.InitApp()

	fmt.Printf("App port : %s\n", AppConfig.GetConfig().Port)
	h.Handlers("/public/", h.PublicHandler)
	h.Handlers("/login", h.LoginHandler)
	h.Handlers("/", h.DefaultHandler)

	//TODO : handle logger error
	log, _ := a.NewFileLogger(c.DefaultLoggerConfig())
	log.Printf("log happened yay!!!")

	dbconf := c.DefaultDBConfig()
	db, _ := a.NewDatabase(dbconf)
	defer db.Database().Close()

	err := db.Database().Ping()

	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	http.ListenAndServe(":"+AppConfig.GetConfig().Port, nil)
	return
}
