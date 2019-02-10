package handler

import (
	"go-member/internal/pkg/member"
	"go-member/internal/pkg/pingpong"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingpong.PingPong).Methods("GET")
	router.HandleFunc("/ping", pingpong.PingPongPost).Methods("POST")
	router.HandleFunc("/member", member.CreateMemberAccount).Methods("POST")
	return router
}
