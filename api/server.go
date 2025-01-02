package api

import (
	db "github.com/WillChen936/simple-bank/db/sqlc"
	"github.com/WillChen936/simple-bank/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	config  utils.Config
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(config utils.Config, queries *db.Queries) *Server {
	server := &Server{
		config:  config,
		queries: queries,
	}

	server.SetupRouter()

	return server
}

func (server *Server) SetupRouter() {
	router := gin.Default()

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
