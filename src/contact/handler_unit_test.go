package contact

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valverde.thiago/go-agenda-api/random"
	"gopkg.in/mgo.v2"
)

type testCase struct {
	name          string
	path          string
	buildRequest  func() interface{}
	buildStubs    func(store *MockStore)
	checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func buildEmptyRequest() interface{} {
	return nil
}

func buildDefaultContactRequest() interface{} {
	return contactRequest{
		Name:   expectedContact.Name,
		Email:  expectedContact.Email,
		Active: true,
	}
}

func TestCreateContactEnpoint(t *testing.T) {
	path := "/contacts"
	testCases := []testCase{
		{
			name:         "OK",
			path:         path,
			buildRequest: buildDefaultContactRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					Create(gomock.Any()).
					Times(1).
					Return(expectedContact, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchContact(t, recorder.Body, expectedContact)
			},
		},
		{
			name:         "Internal Server Error",
			path:         path,
			buildRequest: buildDefaultContactRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					Create(gomock.Any()).
					Times(1).
					Return(expectedContact, mgo.ErrCursor)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Blank Request",
			path: path,
			buildRequest: func() interface{} {
				return contactRequest{}
			},
			buildStubs: func(store *MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	run(t, testCases, http.MethodPost)
}
func TestListContactsEnpoint(t *testing.T) {
	path := "/contacts"
	contactList := createRandomContactList(t, random.RandomName(), 10, false)
	testCases := []testCase{
		{
			name:         "OK",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					GetAll().
					Times(1).
					Return(contactList, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContactList(t, recorder.Body, contactList)
			},
		},
		{
			name:         "Internal Server Error",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					GetAll().
					Times(1).
					Return(nil, mgo.ErrCursor)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:         "Blank Result",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					GetAll().
					Times(1).
					Return(nil, mgo.ErrNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
	}
	run(t, testCases, http.MethodGet)
}
func TestSearchContactsEnpoint(t *testing.T) {
	searchTerm := random.RandomName()
	path := fmt.Sprintf("/contacts?term=%s", searchTerm)
	contactList := createRandomContactList(t, random.RandomName(), 10, false)
	testCases := []testCase{
		{
			name:         "OK",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByName(gomock.Eq(searchTerm)).
					Times(1).
					Return(contactList, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContactList(t, recorder.Body, contactList)
			},
		},
		{
			name:         "Internal Server Error",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByName(gomock.Eq(searchTerm)).
					Times(1).
					Return(nil, mgo.ErrCursor)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:         "Blank Result",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByName(gomock.Eq(searchTerm)).
					Times(1).
					Return(nil, mgo.ErrNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
	}
	run(t, testCases, http.MethodGet)
}

func TestGetContactEndpoint(t *testing.T) {
	contact := createRandomContact(t, random.RandomName(), false)
	path := fmt.Sprintf("/contacts/%s", contact.ID.Hex())

	testCases := []testCase{
		{
			name:         "OK",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByID(gomock.Eq(contact.ID.Hex())).
					Times(1).
					Return(contact, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContact(t, recorder.Body, contact)
			},
		}, {
			name:         "Not Found",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByID(gomock.Eq(contact.ID.Hex())).
					Times(1).
					Return(Contact{}, mgo.ErrNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		}, {
			name:         "Internal Error",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByID(gomock.Eq(contact.ID.Hex())).
					Times(1).
					Return(contact, mgo.ErrCursor)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	run(t, testCases, http.MethodGet)

}

func TestUpdateContactEnpoint(t *testing.T) {
	path := fmt.Sprintf("/contacts/%s", expectedContact.ID.Hex())
	testCases := []testCase{
		{
			name:         "OK",
			path:         path,
			buildRequest: buildDefaultContactRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					Update(gomock.Any()).
					Times(1).
					Return(nil)
				store.EXPECT().
					FindByID(expectedContact.ID.Hex()).
					Times(1).
					Return(expectedContact, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContact(t, recorder.Body, expectedContact)
			},
		},
		{
			name:         "Internal Server Error",
			path:         path,
			buildRequest: buildDefaultContactRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					Update(gomock.Any()).
					Times(1).
					Return(mgo.ErrCursor)
				store.EXPECT().
					FindByID(expectedContact.ID.Hex()).
					Times(1).
					Return(expectedContact, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Blank Request",
			path: path,
			buildRequest: func() interface{} {
				return contactRequest{}
			},
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByID(expectedContact.ID.Hex()).
					Times(1).
					Return(expectedContact, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	run(t, testCases, http.MethodPut)
}

func TestDeleteContactEnpoint(t *testing.T) {
	path := fmt.Sprintf("/contacts/%s", expectedContact.ID.Hex())
	testCases := []testCase{
		{
			name:         "OK",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					Delete(gomock.Eq(expectedContact.ID.Hex())).
					Times(1).
					Return(nil)
				store.EXPECT().
					FindByID(gomock.Eq(expectedContact.ID.Hex())).
					Times(1).
					Return(expectedContact, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:         "Internal Server Error",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					Delete(gomock.Eq(expectedContact.ID.Hex())).
					Times(1).
					Return(mgo.ErrCursor)
				store.EXPECT().
					FindByID(expectedContact.ID.Hex()).
					Times(1).
					Return(expectedContact, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:         "Not Found",
			path:         path,
			buildRequest: buildEmptyRequest,
			buildStubs: func(store *MockStore) {
				store.EXPECT().
					FindByID(expectedContact.ID.Hex()).
					Times(1).
					Return(Contact{}, mgo.ErrNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}
	run(t, testCases, http.MethodDelete)
}

func run(t *testing.T, testCases []testCase, method string) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockStore := NewMockStore(ctrl)

			//build stubs
			testCase.buildStubs(mockStore)
			// start http server and send the request
			server := NewServer(mockStore, gin.Default(), &testConfig)
			recorder := httptest.NewRecorder()
			requestObject := testCase.buildRequest()
			var request *http.Request
			var err error
			if requestObject == nil {
				request, err = http.NewRequest(method, testCase.path, nil)
			} else {
				request, err = http.NewRequest(method, testCase.path,
					sendObjectAsRequestBody(t, requestObject))
			}
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			//check response
			testCase.checkResponse(t, recorder)
		})
	}
}
