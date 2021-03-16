package contact

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/valverde.thiago/go-agenda-api/util"
	"gopkg.in/mgo.v2/bson"
)

func TestCreateContactIntegrationEnpoint(t *testing.T) {
	path := "/contacts"
	testCases := []struct {
		name          string
		buildRequest  func() interface{}
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:         "OK",
			buildRequest: buildDefaultContactRequest,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				returnedContact := readContactFromResponseBody(t, recorder.Body)
				require.NotEmpty(t, returnedContact)
				require.NotEmpty(t, returnedContact.ID)
			},
		},
		{
			name: "Blank Request",
			buildRequest: func() interface{} {
				return contactRequest{}
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			it(func() {
				// start http server and send the request
				server := NewServer(testDbStore, gin.Default(), &testConfig)
				recorder := httptest.NewRecorder()
				requestObject := testCase.buildRequest()
				request, err := http.NewRequest(http.MethodPost, path,
					sendObjectAsRequestBody(t, requestObject))
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, request)
				//check response
				testCase.checkResponse(t, recorder)
			})
		})
	}
}
func TestListContactsIntegrationEnpoint(t *testing.T) {
	path := "/contacts"
	var contactList []Contact

	testCases := []struct {
		name          string
		buildSeed     func(store Store)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildSeed: func(store Store) {
				contactList = createRandomContactList(t, util.RandomName(), 10, true)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContactList(t, recorder.Body, contactList)
			},
		},
		{
			name: "Blank Result",
			buildSeed: func(store Store) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			it(func() {

				//build seeds
				testCase.buildSeed(testDbStore)
				// start http server and send the request
				server := NewServer(testDbStore, gin.Default(), &testConfig)
				recorder := httptest.NewRecorder()
				request, err := http.NewRequest(http.MethodGet, path, nil)
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, request)
				//check response
				testCase.checkResponse(t, recorder)

			})
		})
	}
}

func TestSearchContactsIntegrationEnpoint(t *testing.T) {
	searchTerm := util.RandomName()
	path := fmt.Sprintf("/contacts?term=%s", searchTerm)
	var contactList []Contact
	testCases := []struct {
		name          string
		buildSeed     func(store Store)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildSeed: func(store Store) {
				contactList = createRandomContactList(t, searchTerm, 10, true)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContactList(t, recorder.Body, contactList)
			},
		},
		{
			name: "Blank Result",
			buildSeed: func(store Store) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			it(func() {

				//build seeds
				testCase.buildSeed(testDbStore)
				// start http server and send the request
				server := NewServer(testDbStore, gin.Default(), &testConfig)
				recorder := httptest.NewRecorder()
				request, err := http.NewRequest(http.MethodGet, path, nil)
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, request)
				//check response
				testCase.checkResponse(t, recorder)

			})
		})
	}
}

func TestGetContactIntegrationEndpoint(t *testing.T) {
	var contact Contact
	testCases := []struct {
		name          string
		buildSeed     func(store Store) string
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildSeed: func(store Store) string {
				contact = createRandomContact(t, util.RandomName(), true)
				return fmt.Sprintf("/contacts/%s", contact.ID.Hex())
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContact(t, recorder.Body, contact)
			},
		}, {
			name: "Not Found",
			buildSeed: func(store Store) string {
				contact = createRandomContact(t, util.RandomName(), true)
				return fmt.Sprintf("/contacts/%s", bson.NewObjectId().Hex())
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			it(func() {

				//build seeds
				path := testCase.buildSeed(testDbStore)
				// start http server and send the request
				server := NewServer(testDbStore, gin.Default(), &testConfig)
				recorder := httptest.NewRecorder()
				request, err := http.NewRequest(http.MethodGet, path, nil)
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, request)
				//check response
				testCase.checkResponse(t, recorder)

			})
		})
	}

}

func TestUpdateContactIntegrationEnpoint(t *testing.T) {
	var contact Contact
	testCases := []struct {
		name          string
		buildRequest  func() interface{}
		buildSeed     func(store Store) string
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildRequest: func() interface{} {
				return contactRequest{
					Name:   contact.Name,
					Email:  contact.Email,
					Active: true,
				}
			},
			buildSeed: func(store Store) string {
				contact = createRandomContact(t, util.RandomName(), true)
				return fmt.Sprintf("/contacts/%s", contact.ID.Hex())
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchContact(t, recorder.Body, contact)
			},
		},
		{
			name: "Blank Request",
			buildRequest: func() interface{} {
				return contactRequest{}
			},
			buildSeed: func(store Store) string {
				contact := createRandomContact(t, util.RandomName(), true)
				return fmt.Sprintf("/contacts/%s", contact.ID.Hex())
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			it(func() {

				//build seeds
				path := testCase.buildSeed(testDbStore)
				// start http server and send the request
				server := NewServer(testDbStore, gin.Default(), &testConfig)
				recorder := httptest.NewRecorder()
				request, err := http.NewRequest(http.MethodPut,
					path, sendObjectAsRequestBody(t, testCase.buildRequest()))
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, request)
				//check response
				testCase.checkResponse(t, recorder)

			})
		})
	}
}

func TestDeleteContactIntegrationEnpoint(t *testing.T) {
	testCases := []struct {
		name          string
		buildRequest  func() interface{}
		buildSeed     func() string
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:         "OK",
			buildRequest: buildEmptyRequest,
			buildSeed: func() string {
				contact := createRandomContact(t, util.RandomName(), true)
				return fmt.Sprintf("/contacts/%s", contact.ID.Hex())
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:         "Not Found",
			buildRequest: buildEmptyRequest,
			buildSeed: func() string {
				return fmt.Sprintf("/contacts/%s", bson.NewObjectId().Hex())
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			it(func() {
				//build seeds
				path := testCase.buildSeed()
				// start http server and send the request
				server := NewServer(testDbStore, gin.Default(), &testConfig)
				recorder := httptest.NewRecorder()
				request, err := http.NewRequest(http.MethodDelete,
					path, nil)
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, request)
				//check response
				testCase.checkResponse(t, recorder)

			})
		})
	}
}
