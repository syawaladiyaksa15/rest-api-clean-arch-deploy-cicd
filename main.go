package main

import (
	"rest-api-clean-arch/config"
	"rest-api-clean-arch/factory"
	"rest-api-clean-arch/routes"
)

func main() {
	// connection db
	dbConn := config.InitDB()

	// factory
	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)

	e.Logger.Fatal(e.Start(":8000"))
}
