package server

import (
	"github.com/gin-gonic/gin"
	"github.com/invowk/invowk-cli/internal/issue"
	"github.com/invowk/invowk-cli/internal/tui/bubble/textinput"
	"net/http"
)

type HttpTuiServer struct {
	router *gin.Engine
}

func (server *HttpTuiServer) Start() {
	go func() {
		err := server.router.Run("localhost:8081")
		issue.Handle(err, nil)
	}()

	for {
		_, err := http.Get("http://localhost:8081/health")
		if err == nil {
			break
		}
	}
}

func NewHttpTuiServer() HttpTuiServer {
	router := gin.Default()
	router.GET("/bubble", func(context *gin.Context) {
		textinput.Bubble()
	})

	router.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	return HttpTuiServer{
		router: router,
	}
}
