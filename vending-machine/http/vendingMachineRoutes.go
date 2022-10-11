package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"vending-machine/auth"
)

type VendingMachineRoutes struct {
	userHandler         UserHandler
	productHandler      ProductHandler
	transactionsHandler TransactionsHandler
	authMW              auth.AuthorizationMechanism
}

type VendingMachineRoutesInterface interface {
	SetVendingMachineRoutes(router *mux.Router) *mux.Router
}

func InitVendingMachineRoutes(user UserHandler, product ProductHandler, transactions TransactionsHandler,
	auth auth.AuthorizationMechanism) VendingMachineRoutes {
	return VendingMachineRoutes{
		userHandler:         user,
		productHandler:      product,
		transactionsHandler: transactions,
		authMW:              auth,
	}
}

func (r VendingMachineRoutes) SetVendingMachineRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user", r.userHandler.NewUser).Methods(http.MethodPost)
	router.HandleFunc("/authenticate", r.userHandler.AuthenticateUser).Methods(http.MethodPost)
	router.HandleFunc("/product", r.productHandler.GetProductList).Methods(http.MethodGet)

	protectedBuyer := router.PathPrefix("/user").Subrouter()
	protectedBuyer.Use(r.authMW.BuyerAuthorization)
	protectedBuyer.HandleFunc("", r.userHandler.UserDeposit).Methods(http.MethodPatch)
	protectedBuyer.HandleFunc("", r.userHandler.UserReset).Methods(http.MethodPut)
	protectedBuyer.HandleFunc("/buy", r.transactionsHandler.Purchase).Methods(http.MethodPost)

	protectedSeller := router.PathPrefix("/product").Subrouter()
	protectedSeller.Use(r.authMW.SellerAuthorization)
	protectedSeller.HandleFunc("", r.productHandler.AddProduct).Methods(http.MethodPost)
	protectedSeller.HandleFunc("", r.productHandler.EditProduct).Methods(http.MethodPut)
	protectedSeller.HandleFunc("", r.productHandler.RemoveProduct).Methods(http.MethodDelete)
	return router
}
