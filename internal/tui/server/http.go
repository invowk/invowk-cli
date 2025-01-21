package server

import (
	"github.com/invowk/invowk-cli/internal/tui/bubble/textinput"
	"net/http"
)

type HttpTuiServer struct {
	router *http.ServeMux
}

func (server *HttpTuiServer) Start() error {
	go func() {
		err := http.ListenAndServe(":8081", server.router)
		if err != nil {
			return
		}
	}()

	for {
		_, err := http.Get("http://localhost:8081/invowk-tui/health")
		if err == nil {
			break
		}
	}
	return nil
}

func NewHttpTuiServer() HttpTuiServer {
	router := http.NewServeMux()

	router.HandleFunc("GET /invowk-tui/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("GET /invowk-tui/bubble", func(w http.ResponseWriter, r *http.Request) {
		textinput.Bubble()
	})

	return HttpTuiServer{router: router}
}
