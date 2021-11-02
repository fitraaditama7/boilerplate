package main

import (
	"arka/cmd/config"
	"arka/pkg/database"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	Migrate()
}

func Migrate() {
	config := config.LoadConfig()
	db, err := database.InitDB(config.DBConfig)
	if err != nil {
		logrus.Error(err)
	}

	file, err := os.Open("./migrations/00000_create_table_user.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Create Database
	_, err = db.Exec(string(b))
	if err != nil {
		panic(err)
	}

	file, err = os.Open("./migrations/00001_populate_data_table_user.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err = ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Create Database
	_, err = db.Exec(string(b))
	if err != nil {
		panic(err)
	}
}
