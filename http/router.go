package http

import "github.com/gorilla/mux"

type Router struct {
	companyRoutes RaptorsRoutesInterface
}

func InitMainRouter(r RaptorsRoutesInterface) *Router {
	return &Router{
		raptorsRoutes: r,
	}
}

func (r *Router) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = r.companyRoutes.SetRaptorsRoutes(router)
	return router
}
