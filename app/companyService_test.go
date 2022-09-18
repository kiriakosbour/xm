package app

import (
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
	"xm/domain"
	mock_domain "xm/mocks"
)

func TestCompanyService_CreateCompany_Ok(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	crud := mock_domain.NewMockCompanyCrudInterface(mockCtrl)
	client := mock_domain.NewMockHelperHttpClientInterface(mockCtrl)
	service := CompanyServiceInit(crud, client, &http.Client{})
	mockRecord := &domain.Company{
		Name:    "kyriakos",
		Code:    1234,
		Country: "Us",
		Website: "www.test.com",
		Phone:   60995555,
	}
	expectedResult := "R3dtfDfkJz3"
	crud.EXPECT().SetKey(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	client.EXPECT().IpapiRequest(gomock.Any()).Return(domain.Helper{CountryName: "Cyprus"})
	result, err := service.CreateCompany(*mockRecord, "")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if expectedResult != result {
		t.Errorf("Actual result not equal to expected ")
	}
}

func TestCompanyService_CreateCompany_FailCountry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	crud := mock_domain.NewMockCompanyCrudInterface(mockCtrl)
	client := mock_domain.NewMockHelperHttpClientInterface(mockCtrl)
	service := CompanyServiceInit(crud, client, &http.Client{})
	mockRecord := &domain.Company{
		Name:    "kyriakos",
		Code:    1234,
		Country: "Us",
		Website: "www.test.com",
		Phone:   60995555,
	}
	expectedError := errors.New("unauthorised country")
	client.EXPECT().IpapiRequest(gomock.Any()).Return(domain.Helper{CountryName: "Greece"})
	_, err := service.CreateCompany(*mockRecord, "")
	if err == nil {
		t.Error("Error is expected: ", expectedError)
	}
	if expectedError.Error() != err.Error() {
		t.Errorf("Actual result not equal to expected ")
	}
}
func TestCompanyService_DelCompany_Ok(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	crud := mock_domain.NewMockCompanyCrudInterface(mockCtrl)
	client := mock_domain.NewMockHelperHttpClientInterface(mockCtrl)
	service := CompanyServiceInit(crud, client, &http.Client{})
	id := ""
	client.EXPECT().IpapiRequest(gomock.Any()).Return(domain.Helper{CountryName: "Cyprus"})
	crud.EXPECT().DelKey(gomock.Any()).Return(nil)
	err := service.DelCompany(id)
	if err != nil {
		t.Error("Error is expected: ", err.Error())
	}
	//if expectedResult != result {
	//	t.Errorf("Actual result not equal to expected ")
	//}
}
