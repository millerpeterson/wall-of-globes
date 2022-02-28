package handlers

import (
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"net/http"
	"strconv"
)

func Status(w http.ResponseWriter) {
	_, err := fmt.Fprint(w, "OK")
	if err != nil {
		panic(err)
	}
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
	_, err := fmt.Fprint(w, "OK")
	if err != nil {
		panic(err)
	}
}

func Stop(w http.ResponseWriter, plyr player.Player) {
	plyr.Stop()
	_, err := fmt.Fprint(w, "OK")
	if err != nil {
		panic(err)
	}
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	_, err := fmt.Fprint(w, "Not Found")
	if err != nil {
		panic(err)
	}
}

func BadRequest(w http.ResponseWriter, details string) {
	w.WriteHeader(400)
	_, err := fmt.Fprintf(w, "Bad request: %v", details)
	if err != nil {
		panic(err)
	}
}

func Handler(plyr player.Player) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/status" {
			Status(w)
		} else if r.Method == http.MethodPost && r.URL.Path == "/play" {
			Play(w, r, plyr)
		} else if r.Method == http.MethodPost && r.URL.Path == "/stop" {
			Stop(w, plyr)
		} else {
			NotFound(w)
		}
	}
}
