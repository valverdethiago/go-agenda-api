package contact

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valverde.thiago/go-agenda-api/util"
	"gopkg.in/mgo.v2"
)

func createRandomContact(t *testing.T) Contact {
	arg := Contact{
		Name:   util.RandomName(),
		Email:  util.RandomEmail(),
		Active: true,
	}
	contact, err := testDbStore.Create(arg)
	require.NoError(t, err)
	require.NotEmpty(t, contact)
	require.NotNil(t, contact.ID)
	require.Equal(t, arg.Name, contact.Name)
	require.Equal(t, arg.Email, contact.Email)
	require.Equal(t, arg.Active, contact.Active)
	return contact
}

func createRandomContactList(t *testing.T, size int) []Contact {
	result := make([]Contact, size)
	for i := range result {
		result[i] = createRandomContact(t)
	}
	return result
}

func assertSearchResult(t *testing.T, originalContact Contact, contactList []Contact) {
	require.NotEmpty(t, contactList)
	require.Equal(t, len(contactList), 1)
	require.Equal(t, originalContact.ID, contactList[0].ID)
	require.Equal(t, originalContact.Active, contactList[0].Active)
	require.Equal(t, originalContact.Email, contactList[0].Email)
	require.Equal(t, originalContact.Name, contactList[0].Name)
}

func TestCreateContact(t *testing.T) {
	it(func() {
		createRandomContact(t)
	})
}

func TestGetContact(t *testing.T) {
	it(func() {
		contact := createRandomContact(t)
		dbContact, err := testDbStore.FindByID(contact.ID.Hex())
		require.NoError(t, err)
		require.NotEmpty(t, dbContact)
		require.Equal(t, contact.ID.String(), dbContact.ID.String())
		require.Equal(t, contact.Email, dbContact.Email)
		require.Equal(t, contact.Name, dbContact.Name)
		require.Equal(t, contact.Active, dbContact.Active)
	})
}
func TestUpdateContact(t *testing.T) {
	it(func() {
		contact := createRandomContact(t)
		arg := Contact{
			ID:     contact.ID,
			Name:   util.RandomName(),
			Email:  util.RandomEmail(),
			Active: false,
		}
		err := testDbStore.Update(arg)
		require.NoError(t, err)
		dbContact, err := testDbStore.FindByID(contact.ID.Hex())
		require.NoError(t, err)
		require.NotEmpty(t, dbContact)
		require.Equal(t, contact.ID.String(), dbContact.ID.String())
		require.NotEqual(t, contact.Email, dbContact.Email)
		require.NotEqual(t, contact.Name, dbContact.Name)
		require.NotEqual(t, contact.Active, dbContact.Active)
	})
}
func TestDeleteteContact(t *testing.T) {
	it(func() {
		contact := createRandomContact(t)
		err := testDbStore.Delete(contact.ID.Hex())
		require.NoError(t, err)
		dbContact, err := testDbStore.FindByID(contact.ID.Hex())
		require.Error(t, err, mgo.ErrNotFound)
		require.Empty(t, dbContact)
	})
}

func TestGetAll(t *testing.T) {
	it(func() {
		size := 10
		createRandomContactList(t, size)
		dbContacts, err := testDbStore.GetAll()
		require.NoError(t, err)
		require.NotEmpty(t, dbContacts)
		require.Equal(t, len(dbContacts), size)
	})
}

func TestFindByName(t *testing.T) {
	it(func() {
		firstName := util.RandomName()
		lastName := util.RandomName()
		name := fmt.Sprintf("%s %s", firstName, lastName)
		contact := createRandomContact(t)
		contact.Name = name
		err := testDbStore.Update(contact)
		require.NoError(t, err)
		contactsByFirstName, err := testDbStore.FindByName(firstName)
		require.NoError(t, err)
		assertSearchResult(t, contact, contactsByFirstName)
		contactsByLastName, err := testDbStore.FindByName(lastName)
		require.NoError(t, err)
		assertSearchResult(t, contact, contactsByLastName)
	})
}
