package server

import (
	"github.com/cominotti/invowk/internal/issue"
	"github.com/cominotti/invowk/internal/tui/bubble/textinput"
	"github.com/gin-gonic/gin"
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
