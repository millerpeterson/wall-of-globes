package main

import (
	handlers "github.com/millerpeterson/wall-of-globes/internal/http"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"github.com/millerpeterson/wall-of-globes/internal/vlc"
	"log"
	"net/http"
)

var addr = ":8081"

func main() {
	var plyr player.Player = &vlc.Player{}
	log.Printf("Listening on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(handlers.Handler(plyr))))
}
