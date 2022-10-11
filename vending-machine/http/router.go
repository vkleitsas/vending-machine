package http

import "github.com/gorilla/mux"

type Router struct {
	vendingMachineRoutes VendingMachineRoutesInterface
}

func InitMainRouter(v VendingMachineRoutesInterface) *Router {
	return &Router{
		vendingMachineRoutes: v,
	}
}

func (r *Router) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = r.vendingMachineRoutes.SetVendingMachineRoutes(router)
	return router
}
