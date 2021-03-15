package contact

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	beforeEach "github.com/jknair0/beforeeach"
	"github.com/stretchr/testify/require"
	"github.com/valverde.thiago/go-agenda-api/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var testConfig util.Config
var testDbStore Store
var testDatabase *mgo.Database

var it = beforeEach.Create(setUp, tearDown)
var expectedContact = Contact{
	ID:    bson.NewObjectId(),
	Name:  util.RandomName(),
	Email: util.RandomEmail(),
}

func TestMain(m *testing.M) {
	testConfig = util.LoadEnvConfig("../", "test")
	testDatabase = util.ConnectToDatabase(testConfig)
	testDbStore = NewMongoDbStore(testDatabase)
	os.Exit(m.Run())
}

func setUp() {
	testDbStore.ClearDatabase()
}

func tearDown() {
	testDbStore.ClearDatabase()
}
func requireBodyMatchContact(t *testing.T, body *bytes.Buffer, contact Contact) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var bodyContact Contact
	err = json.Unmarshal(data, &bodyContact)
	require.NoError(t, err)
	require.Equal(t, contact, bodyContact)
}

func requireBodyMatchContactList(t *testing.T, body *bytes.Buffer, contacts []Contact) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var bodyContacts []Contact
	err = json.Unmarshal(data, &bodyContacts)
	require.NoError(t, err)
	for _, contact := range contacts {
		dbContact := findContactInList(bodyContacts, contact)
		require.Equal(t, dbContact, contact)
	}
}

func findContactInList(contacts []Contact, contact Contact) Contact {
	for _, element := range contacts {
		if element.ID == contact.ID {
			return element
		}
	}
	return Contact{}
}

func sendObjectAsRequestBody(t *testing.T, obj interface{}) *bytes.Buffer {
	b, err := json.Marshal(obj)
	require.NoError(t, err)
	return bytes.NewBuffer(b)
}

func createRandomContact(t *testing.T, persistData bool) Contact {
	arg := Contact{
		Name:   util.RandomName(),
		Email:  util.RandomEmail(),
		Active: true,
	}
	var contact Contact
	var err error
	if persistData {
		contact, err = testDbStore.Create(arg)
		require.NoError(t, err)
	} else {
		contact = Contact{
			ID:     bson.NewObjectId(),
			Name:   arg.Name,
			Email:  arg.Email,
			Active: true,
		}
	}
	require.NotEmpty(t, contact)
	require.NotNil(t, contact.ID)
	require.Equal(t, arg.Name, contact.Name)
	require.Equal(t, arg.Email, contact.Email)
	require.Equal(t, arg.Active, contact.Active)
	return contact
}

func createRandomContactList(t *testing.T, size int, persistData bool) []Contact {
	result := make([]Contact, size)
	for i := range result {
		result[i] = createRandomContact(t, persistData)
	}
	return result
}
