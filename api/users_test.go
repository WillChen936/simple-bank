package api

import (
	"testing"

	db "github.com/WillChen936/simple-bank/db/sqlc"
	"github.com/WillChen936/simple-bank/utils"
)

func TestCreateUser(t *testing.T) {

}

func createRandomUser() db.User {
	owner := utils.RandomOwner()

	return db.User{
		Username:       owner,
		HashedPassword: utils.RandomString(8),
		FullName:       owner,
		Email:          utils.RandomEmail(),
	}
}
