package main

import (
	"go-member/internal/src/pingpong"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ping", pingpong.PingPong).Methods("GET")
	router.HandleFunc("/ping", pingpong.PingPongPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8050", router))
}
