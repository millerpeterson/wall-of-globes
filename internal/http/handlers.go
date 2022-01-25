package handlers

import (
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/vlc"
	"net/http"
)

func Status(w http.ResponseWriter) {
	fmt.Fprint(w, "OK")
}

func Play(w http.ResponseWriter, r *http.Request, player *vlc.Player) {
	fmt.Fprint(w, "OK")
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	fmt.Fprint(w, "Not Found")
}

func Server(player *vlc.Player) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/status" {
			Status(w)
		} else if r.Method == http.MethodPost && r.URL.Path == "/play" {
			Play(w, r, player)
		} else {
			NotFound(w)
		}
	}
}
