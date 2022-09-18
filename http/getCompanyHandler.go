package http

import (
	"encoding/json"
	"net/http"
	"xm/app"
)

type GetCompanyHandler struct {
	crudService *app.CompanyService
}

func GetCompanyHandlerInit(crudService *app.CompanyService) *GetCompanyHandler {
	return &GetCompanyHandler{
		crudService: crudService,
	}
}

type GetCompanyHandlerInterface interface {
	GetCompany(w http.ResponseWriter, r *http.Request)
}

func (g *GetCompanyHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Empty request timestamp", http.StatusNoContent)
		return
	}

	result, err := g.crudService.RetrieveCompany(id)
	if err != nil {
		http.Error(w, "No entries found", http.StatusNoContent)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
