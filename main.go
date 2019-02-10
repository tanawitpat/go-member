package main

import (
	"go-member/internal/app"
	"go-member/internal/handler"
	"log"
	"net/http"
)

func main() {
	if err := app.InitConfig(); err != nil {
		panic(err)
	}
	log.Println("Initial config: ", app.CFG)

	if err := app.InitErrorMessage(); err != nil {
		panic(err)
	}
	log.Println("Initial error message: ", app.EM)

	router := handler.NewRouter()
	log.Fatal(http.ListenAndServe(":8050", router))
}
