package contact

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/mgo.v2"
)

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
		createRandomContact(t, randomizer.RandomName(), true)
	})
}

func TestGetContact(t *testing.T) {
	it(func() {
		contact := createRandomContact(t, randomizer.RandomName(), true)
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
		contact := createRandomContact(t, randomizer.RandomName(), true)
		arg := Contact{
			ID:     contact.ID,
			Name:   randomizer.RandomName(),
			Email:  randomizer.RandomEmail(),
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
		contact := createRandomContact(t, randomizer.RandomName(), true)
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
		createRandomContactList(t, randomizer.RandomName(), size, true)
		dbContacts, err := testDbStore.GetAll()
		require.NoError(t, err)
		require.NotEmpty(t, dbContacts)
		require.Equal(t, len(dbContacts), size)
	})
}

func TestFindByName(t *testing.T) {
	it(func() {
		firstName := randomizer.RandomName()
		lastName := randomizer.RandomName()
		name := fmt.Sprintf("%s %s", firstName, lastName)
		contact := createRandomContact(t, randomizer.RandomName(), true)
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
