package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/pkg"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(pkg.ErrLoadingENV.Error())
	}

	db, err := pkg.ConnectToDatabase()
	if err != nil {
		log.Fatal(pkg.ErrConnectionDB.Error())
	}

	migrator := db.Migrator()

	var chart models.Chart
	var user models.User
	var node models.Node
	var edge models.Edge

	err = migrator.AutoMigrate(&chart, &user, &node, &edge)
	if err != nil {
		log.Fatalf("%w: %v", pkg.ErrMigratingDB.Error(), err)
	}

	log.Print("auto migrations ran successfully")
}
