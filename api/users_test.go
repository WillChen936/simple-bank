package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockdb "github.com/WillChen936/simple-bank/db/mock"
	db "github.com/WillChen936/simple-bank/db/sqlc"
	"github.com/WillChen936/simple-bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateUserAPI(t *testing.T) {
	user, password := createRandomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username:       user.Username,
					HashedPassword: user.HashedPassword,
					FullName:       user.FullName,
					Email:          user.Email,
				}
				store.EXPECT().CreateUser(gomock.Any(), arg).Times(1).Return(user, nil)
			},
			checkResponse: func(recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := newTestServer(store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(*recorder)
		})
	}
}

func createRandomUser(t *testing.T) (user db.User, password string) {
	owner := utils.RandomOwner()
	username := strings.Fields(owner)[0]
	password = utils.RandomString(8)

	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       username,
		HashedPassword: hashedPassword,
		FullName:       owner,
		Email:          utils.RandomEmail(),
	}

	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
