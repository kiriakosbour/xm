package http

import (
	"encoding/json"
	"net/http"
	"xm/app"
	"xm/domain"
)

type JwtClaimsHandler struct {
	service *app.JwtService
}
type JwtClaimsHandlerInterface interface {
	UrlAddController(w http.ResponseWriter, r *http.Request)
}

func JwtClaimsHandlerInit(claims *app.JwtService) *JwtClaimsHandler {
	return &JwtClaimsHandler{service: claims}
}
func (j *JwtClaimsHandler) JwtTokenController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user domain.User
	errorMsg := json.NewDecoder(r.Body).Decode(&user)
	if errorMsg != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMsg)
	}
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Empty username or password is not allowed")
		return
	}
	result, err := j.service.CreateJwtTokenService(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("There is an error during creation of the token")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
}
