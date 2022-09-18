package app

import (
	"errors"
	"log"
	"net/http"
	"time"
	"xm/domain"
)

type CompanyService struct {
	DB            domain.CompanyCrudInterface
	HelperClient  domain.HelperHttpClientInterface
	HttpClientImp *http.Client
}

func CompanyServiceInit(database domain.CompanyCrudInterface, client domain.HelperHttpClientInterface, clientinterface *http.Client) *CompanyService {
	return &CompanyService{
		DB:            database,
		HelperClient:  client,
		HttpClientImp: clientinterface,
	}
}
func (c *CompanyService) CreateCompany(company domain.Company, id string) (string, error) {
	newComp := domain.Company{}
	if id == "" {
		location := c.HelperClient.IpapiRequest(c.HttpClientImp)
		if location.CountryName != "Cyprus" {
			log.Printf("Only calls from Cyprus are allowed the call was from %s", location.CountryName)
			return "", errors.New("unauthorised country")
		}
		id = newComp.GenerateId(company)
	}
	err := c.DB.SetKey(company, id, time.Minute*60)
	if err != nil {
		log.Printf("Redis set responded with an error %s", err.Error())
		return "", err
	}
	return id, nil

}
func (c *CompanyService) RetrieveCompany(id string) (domain.Company, error) {

	cmp, err := c.DB.GetKey(id)
	newComp := domain.Company{}
	if err != nil {
		log.Printf("Redis set responded with an error %s", err.Error())
		return newComp, err
	}
	return cmp, nil

}
func (c *CompanyService) DelCompany(id string) error {
	location := c.HelperClient.IpapiRequest(c.HttpClientImp)
	if location.CountryName != "Cyprus" {
		log.Printf("Only calls from Cyprus are allowed the call was from %s", location.CountryName)
		return errors.New("unauthorised country")
	}
	err := c.DB.DelKey(id)
	if err != nil {
		log.Printf("An error has occured in deleting the user in Redis %s ", err)
		return err
	}
	return nil
}
