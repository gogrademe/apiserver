package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	h "github.com/gogrademe/apiserver/handlers"
	"github.com/gogrademe/apiserver/store"
	"github.com/mattaitchison/envconfig"

	"github.com/gin-gonic/gin"
)

var version string

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var listenAddress = envconfig.String("listen_address", ":5005", "address to listen on")
var dbName = envconfig.String("db_name", "dev_go_grade", "rethinkdb database name")
var dbAddress = envconfig.String("RETHINKDB_PORT_28015_TCP", "localhost:28015", "rethinkdb address")
var bootstrap = envconfig.Bool("create_tables", false, "create tables in db")
var testData = envconfig.Bool("test_data", false, "insert test data into db")

func main() {
	log.Println("Starting api server Version:", version)

	log.Println(bootstrap, testData)

	// FIXME: I don't think this is needed any more.
	// I think it was only for wercker.
	// Actually I think this may have been for docker links.
	dbAddress = strings.Trim(dbAddress, "tcp://")

	if err := store.Connect(dbAddress, dbName); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	store.SetupDB(bootstrap, testData)

	r := gin.Default()

	h.SetupHandlers(r)

	r.Run(listenAddress)
}
