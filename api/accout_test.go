package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	mockdb "github.com/tetrex/backend-masterclass-go/db/mock"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/util"
	"go.uber.org/mock/gomock"
)

func TestGetAcountApi(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	// build stubs
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// start http server
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/v1/account/%d", account.ID)

	request := httptest.NewRequest(http.MethodGet, url, nil)

	e := echo.New()
	e.GET("/v1/account/:id", server.getAccount)
	e.ServeHTTP(recorder, request)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomAccount() db.Account {
	return db.Account{
		Owner:    util.RandomOwner(),
		Currency: util.RandomCurrency(),
		ID:       util.RandomInt(1, 1000),
		Balance:  util.RandomMoney(),
	}
}
