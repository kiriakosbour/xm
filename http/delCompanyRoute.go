package http

import (
	"net/http"
	"xm/app"
)

type DelCompanyHandler struct {
	crudService *app.CompanyService
}

func DelCompanyHandlerInit(crudService *app.CompanyService) *DelCompanyHandler {
	return &DelCompanyHandler{
		crudService: crudService,
	}
}

type DelCompanyHandlerInterface interface {
	DelCompany(w http.ResponseWriter, r *http.Request)
}

func (g *DelCompanyHandler) DelCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Empty id", http.StatusNoContent)
		return
	}

	err := g.crudService.DelCompany(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
