package main

import (
	"go-member/internal/app"
	"go-member/internal/handler"
	"net/http"
)

func main() {
	log := app.InitLogger()

	if err := app.InitConfig(); err != nil {
		panic(err)
	}
	log.Infof("Initial config: %+v", app.CFG)

	if err := app.InitErrorMessage(); err != nil {
		panic(err)
	}
	log.Infof("Initial error message: %+v", app.EM)

	router := handler.NewRouter()
	log.Fatal(http.ListenAndServe(":8050", router))
}
