package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"xm/domain"
)

type HelperHttpClient struct {
}

func HelperHttpClientInit() *HelperHttpClient {
	return &HelperHttpClient{}
}

func (o *HelperHttpClient) IpapiRequest(client *http.Client) domain.Helper {
	helper := domain.Helper{}
	req, err := http.NewRequest("GET", "https://ipapi.co/8.8.8.8/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error during encoding helper response %s", err)
		return helper
	}
	json.Unmarshal(body, &helper)
	return helper
}
