package handlers

import (
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"net/http"
	"strconv"
)

func Status(w http.ResponseWriter) {
	fmt.Fprint(w, "OK")
}

func Play(w http.ResponseWriter, r *http.Request, plyr player.Player) {
	filePath := r.FormValue("file")
	if filePath == "" {
		BadRequest(w, "No file specified")
		return
	}

	var cropArgs = player.Args{}
	cropArgs.Top, _ = strconv.Atoi(r.FormValue("top"))
	cropArgs.Bottom, _ = strconv.Atoi(r.FormValue("bottom"))
	cropArgs.Left, _ = strconv.Atoi(r.FormValue("left"))
	cropArgs.Right, _ = strconv.Atoi(r.FormValue("right"))

	plyr.Play(filePath, cropArgs)
	fmt.Fprint(w, "OK")
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	fmt.Fprint(w, "Not Found")
}

func BadRequest(w http.ResponseWriter, details string) {
	w.WriteHeader(400)
	fmt.Fprintf(w, "Bad request: %v", details)
}

func Server(player player.Player) func(http.ResponseWriter, *http.Request) {
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
