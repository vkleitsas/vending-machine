package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vending-machine/app"
	"vending-machine/auth"
	"vending-machine/db"
	h "vending-machine/http"
)

func main() {
	db.InitDB()
	dbClient, err := db.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Unable to create DB client %s", err)
	}

	postgres := db.NewPostgresClient(dbClient)
	repository := db.NewSQLStore()
	authMechanism := auth.NewAuthMechanism()
	userCrud := app.NewUserCrudService(postgres, repository, authMechanism)
	productCrud := app.NewProductCrudService(postgres, repository)
	transactionsCrud := app.NewTransactionsCrudService(postgres, repository)
	userHandler := h.NewUserHandler(*userCrud)
	productHandler := h.NewProductHandler(*productCrud)
	transactionsHandler := h.NewTransactionsHandler(*transactionsCrud)
	authMiddleware := auth.NewAuthorizationMiddleware(authMechanism)
	vendingMachineRoutes := h.InitVendingMachineRoutes(*userHandler, *productHandler, *transactionsHandler, authMiddleware)
	routes := h.InitMainRouter(vendingMachineRoutes)
	muxRouter := routes.InitRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), muxRouter))
}
