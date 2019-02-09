package main

import (
	"go-member/internal/app"
	"go-member/internal/pkg/member"
	"go-member/internal/pkg/pingpong"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if err := app.InitConfig(); err != nil {
		panic(err)
	}
	log.Println("Initial Config: ", app.CFG)

	router := mux.NewRouter()
	router.HandleFunc("/ping", pingpong.PingPong).Methods("GET")
	router.HandleFunc("/ping", pingpong.PingPongPost).Methods("POST")
	router.HandleFunc("/member", member.CreateMemberAccount).Methods("POST")

	log.Fatal(http.ListenAndServe(":8050", router))
}
