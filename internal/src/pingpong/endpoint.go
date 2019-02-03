package pingpong

import (
	"net/http"

	"github.com/go-chi/render"
)

func PingPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func PingPongPost(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, "pong")
	return
}
