package api

import (
	db "github.com/WillChen936/simple-bank/db/sqlc"
	"github.com/WillChen936/simple-bank/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	config utils.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config utils.Config, store db.Store) *Server {
	server := &Server{
		config: config,
		store:  store,
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
