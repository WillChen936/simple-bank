package api

import (
	"os"
	"testing"

	db "github.com/WillChen936/simple-bank/db/sqlc"
	"github.com/WillChen936/simple-bank/utils"
	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{}
	server := NewServer(config, store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
