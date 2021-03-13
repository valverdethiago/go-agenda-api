package contact

import (
	"os"
	"testing"

	beforeEach "github.com/jknair0/beforeeach"
	"github.com/valverde.thiago/go-agenda-api/util"
	"gopkg.in/mgo.v2"
)

var mockStore Store
var testDbStore Store
var testDatabase *mgo.Database

var it = beforeEach.Create(setUp, tearDown)

func TestMain(m *testing.M) {
	config := util.LoadEnvConfig("../", "test")
	testDatabase = util.ConnectToDatabase(config)
	testDbStore = NewMongoDbStore(testDatabase)
	os.Exit(m.Run())
}

func setUp() {
	testDbStore.ClearDatabase()
}

func tearDown() {
	testDbStore.ClearDatabase()
}
