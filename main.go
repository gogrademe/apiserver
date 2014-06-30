package main

import (
	"flag"

	h "github.com/Lanciv/GoGradeAPI/handlers"
	"github.com/Lanciv/GoGradeAPI/store"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/meatballhat/negroni-logrus"
)

var log = logrus.New()

var (
	apiPort  string
	address  string
	dbName   string
	testData bool
)

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) // default
}

func main() {
	flag.StringVar(&apiPort, "apiPort", ":5005", "")
	flag.StringVar(&address, "dbAddress", "localhost:28015", "")
	flag.StringVar(&dbName, "dbName", "dev_go_grade", "")
	flag.BoolVar(&testData, "testData", true, "")
	flag.Parse()

	if err := store.Connect(address, dbName); err != nil {
		log.Fatal("Error setting up database: ", err)
	}

	store.SetupDB(testData)

	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
	n.Use(negroni.HandlerFunc(h.CORSMiddleware))
	n.UseHandler(h.SetupHandlers())

	n.Run(apiPort)

}
