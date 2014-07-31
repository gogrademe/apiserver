package main

import (
	"flag"

	h "github.com/Lanciv/GoGradeAPI/handlers"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var log = logrus.New()

var (
	listenAddr     string
	address        string
	dbName         string
	staticDir      string
	insertTestData bool
)

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) // default
}

func main() {
	flag.StringVar(&listenAddr, "listenAddr", ":5005", "")
	flag.StringVar(&address, "dbAddress", "localhost:28015", "")
	flag.StringVar(&dbName, "dbName", "dev_go_grade", "")
	flag.StringVar(&staticDir, "staticDir", "public", "")
	flag.BoolVar(&insertTestData, "insertTestData", true, "")

	flag.Parse()

	if err := store.Connect(address, dbName); err != nil {
		log.Fatal("Error setting up database: ", err)
	}

	store.SetupDB(insertTestData)

	r := gin.Default()
	r.Static("/app", staticDir)

	h.SetupHandlers(r)

	r.Run(listenAddr)

}
