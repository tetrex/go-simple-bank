package api

import (
	"testing"

	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/util"
	"go.uber.org/mock/gomock"
)

func TestGetAcountApi(t *testing.T) {
	// account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// store := mockdb.NewMockStore(ctrl)

	// // build stubs
	// store.EXPECT().
	// 	GetAccount(gomock.Any(), gomock.Eq(account.ID)).
	// 	Times(1).
	// 	Return(account, nil)

	// // start http server
	// server := NewServer(store)
	// recorder := httptest.NewRecorder()

	// url := fmt.Sprintf("/v1/account/%d", account.ID)
	// fmt.Printf("url::%s", url)

	// request := httptest.NewRequest(http.MethodPost, url, nil)

	// server.router.ServeHTTP(recorder, request)

	// // check response
	// fmt.Println(recorder)
	// require.Equal(t, http.StatusOK, recorder.Code)
}

func randomAccount() db.Account {
	return db.Account{
		Owner:    util.RandomOwner(),
		Currency: util.RandomCurrency(),
		ID:       util.RandomInt(1, 1000),
		Balance:  util.RandomMoney(),
	}
}
