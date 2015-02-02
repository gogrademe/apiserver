package main

import (
	"flag"
	"log"
	"os"
	"strings"

	h "github.com/gogrademe/apiserver/handlers"
	"github.com/gogrademe/apiserver/store"

	"github.com/gin-gonic/gin"
)

var version string

// Borrowed from:
// github.com/progrium/logspout/blob/master/logspout.go#L33
func getopt(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}

func main() {
	log.Println("Starting api server Version:", version)

	port := getopt("PORT", "5005")
	dbName := getopt("DB_NAME", "dev_go_grade")
	bootstrap := getopt("BOOTSTRAP_DB", "") != ""
	testData := getopt("INSERT_TEST_DATA", "") != ""

	// FIXME: I don't think this is needed any more.
	// I think it was only for wercker.
	// Actually I think this may have been for docker links.
	dbAddress := os.Getenv("RETHINKDB_PORT_28015_TCP")
	dbAddress = strings.Trim(dbAddress, "tcp://")

	if dbAddress == "" {
		dbAddress = "localhost:28015"
	}

	flag.Parse()

	if err := store.Connect(dbAddress, dbName); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	store.SetupDB(bootstrap, testData)

	r := gin.Default()

	h.SetupHandlers(r)

	r.Run(":" + port)

}
