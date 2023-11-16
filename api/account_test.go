package api

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/golang/mock/gomock"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	"github.com/techschool/simplebank/api/schema"
// 	mockdb "github.com/techschool/simplebank/db/mock"
// 	db "github.com/techschool/simplebank/db/sqlc"
// 	"github.com/techschool/simplebank/uti"
// )

// func TestGetAccountAPI(t *testing.T) {
// 	account := RandomAccount()

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	testCases := []struct {
// 		name          string
// 		accountID     interface{}
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:      "OK",
// 			accountID: account.ID,
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(account, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchAccount(t, recorder.Body, account)
// 			},
// 		},
// 		{
// 			name:      "Not Found",
// 			accountID: account.ID,
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), account.ID).
// 					Times(1).
// 					Return(account, sql.ErrNoRows)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "Query Param is not int64",
// 			accountID: 0.2,
// 			buildStubs: func(store *mockdb.MockStore) {
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "No connect db",
// 			accountID: account.ID,
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), account.ID).
// 					Times(1).
// 					Return(account, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		// TODO add more test case
// 	}
// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			var url string
// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			// start test server and send reqeust
// 			server := NewServer(store)
// 			recorder := httptest.NewRecorder()
// 			if tc.name == "Query Param is not int64" {
// 				url = fmt.Sprintf("/account/%f", tc.accountID)
// 			} else {
// 				url = fmt.Sprintf("/account/%d", tc.accountID)
// 			}

// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			server.route.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})

// 	}
// }

// func RandomAccount() db.Account {
// 	return db.Account{
// 		ID:       uti.RandomInt(1, 1000),
// 		Owner:    uti.RandomOwner(),
// 		Balance:  uti.RandomMoney(),
// 		Currency: "VND",
// 	}
// }

// func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
// 	data, err := io.ReadAll(body)
// 	require.NoError(t, err)

// 	var response schema.Response
// 	var gotAccount db.Account
// 	err = json.Unmarshal(data, &response)
// 	require.NoError(t, err)
// 	data, err = json.Marshal(response.Data)
// 	require.NoError(t, err)
// 	err = json.Unmarshal(data, &gotAccount)
// 	require.NoError(t, err)
// 	require.Equal(t, account, gotAccount)
// }

// func TestPostAccountAPI(t *testing.T) {
// 	account := RandomAccount()
// 	createAccountParams := getCreateAccountParams(account)
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(store *mockdb.MockStore)
// 		body          CreateAccountRequest
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: RandomCreateAccountRequestOK(createAccountParams),
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					CreateAccount(gomock.Any(), createAccountParams).
// 					Times(1).
// 					Return(account, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchAccount(t, recorder.Body, account)
// 			},
// 		},
// 		{
// 			name: "Missing Balance",
// 			body: RandomCreateAccountRequestMissingBalance(createAccountParams),
// 			buildStubs: func(store *mockdb.MockStore) {
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "Missing Request Time",
// 			body: RandomCreateAccountRequestMissingRequestTime(createAccountParams),
// 			buildStubs: func(store *mockdb.MockStore) {
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "DB not connect",
// 			body: RandomCreateAccountRequestOK(createAccountParams),
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					CreateAccount(gomock.Any(), createAccountParams).
// 					Times(1).
// 					Return(account, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}
// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			var url string
// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			// start test server and send reqeust
// 			server := NewServer(store)
// 			recorder := httptest.NewRecorder()
// 			url = "/accounts/"
// 			var buf bytes.Buffer
// 			err := json.NewEncoder(&buf).Encode(tc.body)
// 			require.NoError(t, err)
// 			request, err := http.NewRequest(http.MethodPost, url, &buf)
// 			require.NoError(t, err)

// 			server.route.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})

// 	}
// }

// func RandomCreateAccountRequestOK(account db.CreateAccountParams) (createAccountRequest CreateAccountRequest) {
// 	createAccountRequest.RequestInfo.RequestId = uuid.New().String()
// 	createAccountRequest.RequestInfo.RequestTime = time.Now().Format(uti.DateTimeLayout)
// 	createAccountRequest.RequestInfo.LangCode = "en"
// 	createAccountRequest.Balance = account.Balance
// 	createAccountRequest.Owner = account.Owner

// 	return
// }

// func RandomCreateAccountRequestMissingBalance(account db.CreateAccountParams) (createAccountRequest CreateAccountRequest) {
// 	createAccountRequest.RequestInfo.RequestId = uuid.New().String()
// 	createAccountRequest.RequestInfo.RequestTime = time.Now().Format(uti.DateTimeLayout)
// 	createAccountRequest.RequestInfo.LangCode = "en"
// 	createAccountRequest.Owner = account.Owner

// 	return
// }

// func RandomCreateAccountRequestMissingRequestTime(account db.CreateAccountParams) (createAccountRequest CreateAccountRequest) {
// 	createAccountRequest.RequestInfo.RequestId = uuid.New().String()
// 	createAccountRequest.RequestInfo.LangCode = "en"
// 	createAccountRequest.Balance = account.Balance
// 	createAccountRequest.Owner = account.Owner

// 	return
// }

// func getCreateAccountParams(account db.Account) (createAccountParams db.CreateAccountParams) {
// 	createAccountParams.Balance = account.Balance
// 	createAccountParams.Currency = "VND"
// 	createAccountParams.Owner = account.Owner
// 	return
// }

// func TestGetListAccount(t *testing.T) {
// 	listAccount := randomListAccount()
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(store *mockdb.MockStore)
// 		url           string
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK List not query parm",
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					ListAccount(gomock.Any(),
// 						db.ListAccountParams{
// 							Offset: 0,
// 							Limit:  10,
// 						}).
// 					Return(listAccount, nil)
// 			},
// 			url: "/accounts/",
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "OK List have query parm",
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					ListAccount(gomock.Any(),
// 						db.ListAccountParams{
// 							Offset: 0,
// 							Limit:  10,
// 						}).
// 					Return(listAccount, nil)
// 			},
// 			url: "/accounts/?page_id=1&page_size=10",
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "DB ERROR",
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					ListAccount(gomock.Any(),
// 						db.ListAccountParams{
// 							Offset: 0,
// 							Limit:  10,
// 						}).
// 					Return(listAccount, sql.ErrConnDone)
// 			},
// 			url: "/accounts/?page_id=1&page_size=10",
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			var url string
// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			// start test server and send reqeust
// 			server := NewServer(store)
// 			recorder := httptest.NewRecorder()
// 			url = tc.url
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			server.route.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})

// 	}
// }

// func randomListAccount() []db.Account {
// 	element := 10
// 	var listAccount []db.Account
// 	for i := 0; i < element; i++ {
// 		listAccount = append(listAccount, RandomAccount())
// 	}
// 	return listAccount
// }
