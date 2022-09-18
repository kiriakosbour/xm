package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CompanyRoutes struct {
	getCompanyHandler    GetTotalRaptorsHandler
	updateCompanyHandler UpdateTotalRaptorsHandler
}

type RaptorsRoutesInterface interface {
	SetRaptorsRoutes(router *mux.Router) *mux.Router
}

func InitRaptorsRoutes(get GetTotalRaptorsHandler, update UpdateTotalRaptorsHandler) CompanyRoutes {
	return CompanyRoutes{
		getCompanyHandler:    get,
		updateCompanyHandler: update,
	}
}
func (a CompanyRoutes) SetRaptorsRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/raptors", a.getCompanyHandler.GetTotalRaptors).Methods(http.MethodGet)
	router.HandleFunc("/api/raptors", a.updateCompanyHandler.UpdateTotalRaptors).Methods(http.MethodPut)
	return router
}
