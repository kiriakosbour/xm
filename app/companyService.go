package app

import "xm/domain"

type ComanyService struct {
	DB           domain.CompanyCrudInterface
	HelperClient domain.HelperHttpClientInterface
}
