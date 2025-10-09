package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
	routes(r *gin.Engine)
}

type GinServer struct {
	router *gin.Engine
	Client
}

func NewServer(db Client) Server {
	server := &GinServer{
		gin.Default(),
		db,
	}
	server.routes(server.router)
	return server
}

func (s *GinServer) Start() error {
	slog.Info("serving at port 8372")
	return s.router.Run(":8372")
}