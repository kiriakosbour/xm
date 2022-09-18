package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
	"xm/app"
	"xm/auth"
	"xm/db"
	xmhttp "xm/http"
)

func main() {
	router := mux.NewRouter()
	jwtRepo := auth.AuthMechInit()
	jwtService := app.JwtServiceInit(jwtRepo)
	redisRepo := db.CompanyCrudRepoInit()
	client := &http.Client{Timeout: 10 * time.Second}
	helper := xmhttp.HelperHttpClientInit()
	cmpService := app.CompanyServiceInit(redisRepo, helper, client)
	getRoute := xmhttp.GetCompanyHandlerInit(cmpService)
	setRoute := xmhttp.SetCompanyHandlerInit(cmpService)
	delRoute := xmhttp.DelCompanyHandlerInit(cmpService)
	jwtRoute := xmhttp.JwtClaimsHandlerInit(jwtService)
	authMw := xmhttp.AuthenticationMwInit(jwtRepo)
	router.HandleFunc("/auth", jwtRoute.JwtTokenController).Methods(http.MethodPost)
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(authMw.AuthenticationMW)
	router.HandleFunc("/company/", setRoute.SetCompany).Methods(http.MethodPost)
	protected.HandleFunc("/company/", setRoute.SetCompany).Methods(http.MethodPut)
	protected.HandleFunc("/company/", delRoute.DelCompany).Methods(http.MethodDelete)
	protected.HandleFunc("/company/", getRoute.GetCompany).Methods(http.MethodGet)
	//address := os.Getenv("SERVER_ADDR")
	address := "0.0.0.0"
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("No port given. Defaulting to port %s\n", port)
	}
	log.Print(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
