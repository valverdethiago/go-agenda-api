package contact

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	contactCollectionName = "contacts"
)

// Store interface to interact with contact colletion
type Store interface {
	GetAll() ([]Contact, error)
	FindByID(ID string) (Contact, error)
	FindByName(name string) ([]Contact, error)
	Create(contact Contact) (Contact, error)
	Update(contact Contact) error
	Delete(ID string) error
	ClearDatabase() error
}

// MongoDbStore  contact store implementaion for mongodb
type MongoDbStore struct {
	database   *mgo.Database
	collection *mgo.Collection
}

// NewMongoDbStore builds a new instance of store
func NewMongoDbStore(database *mgo.Database) Store {
	store := &MongoDbStore{
		database: database,
	}
	store.connect(contactCollectionName)
	return store
}

func (store *MongoDbStore) connect(collectionName string) {
	store.collection = store.database.C(collectionName)
}

// GetAll returns all contacts on the database
func (store *MongoDbStore) GetAll() ([]Contact, error) {
	var contacts []Contact
	err := store.collection.Find(bson.M{}).All(&contacts)
	return contacts, err
}

// FindByID finds the contact with the given id
func (store *MongoDbStore) FindByID(ID string) (Contact, error) {
	var contact Contact
	err := store.collection.FindId(bson.ObjectIdHex(ID)).One(&contact)
	return contact, err
}

// FindByName finds the contact by name
func (store *MongoDbStore) FindByName(name string) ([]Contact, error) {
	var contacts []Contact
	err := store.collection.Find(
		bson.M{"name": bson.RegEx{
			Pattern: name,
			Options: "i",
		},
		},
	).All(&contacts)
	return contacts, err
}

// Create creates a new contact
func (store *MongoDbStore) Create(contact Contact) (Contact, error) {
	contact.ID = bson.NewObjectId()
	err := store.collection.Insert(&contact)
	return contact, err
}

// Delete deletes the contact by id
func (store *MongoDbStore) Delete(id string) error {
	err := store.collection.RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update updates contact
func (store *MongoDbStore) Update(contact Contact) error {
	id := bson.ObjectIdHex(contact.ID.Hex())
	err := store.collection.UpdateId(id, &contact)
	return err
}

func (store *MongoDbStore) ClearDatabase() error {
	return store.collection.DropCollection()
}
