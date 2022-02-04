package handlers

import (
	"bytes"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestStatus(t *testing.T) {
	t.Run("Status endpoint", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/status", nil)
		response := httptest.NewRecorder()

		var logger player.Player = &player.PlayCmdLogger{}
		Handler(logger)(response, request)

		if response.Code != 200 {
			t.Errorf("Unexpected response code %v", response.Code)
		}

		got := response.Body.String()
		want := "OK"

		if got != want {
			t.Errorf("Body is %q, expected %q", got, want)
		}
	})
}

func TestPlay(t *testing.T) {
	t.Run("Play endpoint", func(t *testing.T) {
		postData := url.Values{}
		postData.Set("file", "udp://@225.0.0.1")
		postData.Set("top", "100")
		postData.Set("bottom", "200")
		postData.Set("left", "300")
		postData.Set("right", "400")

		request, _ := http.NewRequest(http.MethodPost, "/play", bytes.NewBufferString(postData.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		response := httptest.NewRecorder()
		var logger player.Player = &player.PlayCmdLogger{}
		Handler(logger)(response, request)

		if response.Code != 200 {
			t.Errorf("Unexpected response code %v", response.Code)
		}

		args := player.Args{Top: 100, Bottom: 200, Left: 300, Right: 400}
		var expectedPlayCmd = player.PlayCmd{File: "udp://@225.0.0.1", Args: args}
		loggedCmd := player.PlayCmd{}
		if len(logger.(*player.PlayCmdLogger).Log) > 0 {
			loggedCmd = logger.(*player.PlayCmdLogger).Log[0]
		}
		if loggedCmd != expectedPlayCmd {
			t.Errorf("Expected play cmd %v, received %q", expectedPlayCmd, loggedCmd)
		}
	})
}

func TestInvalidRoute(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/invalid", nil)
		response := httptest.NewRecorder()

		var logger player.Player = &player.PlayCmdLogger{}
		Handler(logger)(response, request)

		if response.Code != 404 {
			t.Errorf("Unexpected response code %v", response.Code)
		}

		got := response.Body.String()
		want := "Not Found"

		if got != want {
			t.Errorf("Body is %q, expected %q", got, want)
		}
	})
}
