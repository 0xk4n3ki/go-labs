package main

import (
	"github.com/gin-gonic/gin"
)

func (s *GinServer) routes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user", s.AddUser())
	incomingRoutes.GET("/users", s.ListUsers())
	incomingRoutes.PUT("/user/:id", s.UpdateUser())
	incomingRoutes.DELETE("/user/:id", s.DeleteUser())
}