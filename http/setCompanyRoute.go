package http

import (
	"encoding/json"
	"log"
	"net/http"
	"xm/app"
	"xm/domain"
)

type SetCompanyHandler struct {
	crudService *app.CompanyService
}

func SetCompanyHandlerInit(crudService *app.CompanyService) *SetCompanyHandler {
	return &SetCompanyHandler{
		crudService: crudService,
	}
}

type SetCompanyHandlerInterface interface {
	SetCompany(w http.ResponseWriter, r *http.Request)
}

func (g *SetCompanyHandler) SetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := domain.Company{}
	id := r.URL.Query().Get("id")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding the body %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := ""
	if id == "" {
		result, err = g.crudService.CreateCompany(req, "")
	}
	result, err = g.crudService.CreateCompany(req, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
