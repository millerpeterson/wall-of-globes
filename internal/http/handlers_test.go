package handlers

import (
	"github.com/millerpeterson/wall-of-globes/internal/vlc"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	t.Run("Status endpoint", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/status", nil)
		response := httptest.NewRecorder()

		Server(vlc.Logger())(response, request)

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
		request, _ := http.NewRequest(http.MethodPost, "/play", nil)
		response := httptest.NewRecorder()

		logger := vlc.Logger()
		Server(logger)(response, request)

		if response.Code != 200 {
			t.Errorf("Unexpected response code %v", response.Code)
		}

		//expectedPlayCmd := []string{"-vvv", "udp://@225.0.0.1", }
		//if logger {
		//	t.Errorf("Body is %q, expected %q", got, want)
		//}
	})
}

func TestInvalidRoute(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/invalid", nil)
		response := httptest.NewRecorder()

		Server(vlc.Logger())(response, request)

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
