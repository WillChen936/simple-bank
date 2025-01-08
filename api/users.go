package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (server *Server) getUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
