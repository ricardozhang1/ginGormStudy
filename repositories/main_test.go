package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"irisStudy/conf"
	"log"
	"os"
	"testing"
)

var testStore Store
var testDB *gorm.DB

func TestMain(m *testing.M) {
	config, err := conf.LoadConfig("../conf")
	if err != nil {
		log.Fatalf("failed to read config, err: %v", err)
	}

	testDB, err = gorm.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("failed to connect to DB, err: %v", err)
	}

	testStore = NewStore(testDB, config)
	os.Exit(m.Run())
}





