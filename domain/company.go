package domain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/itchyny/base58-go"
	"log"
	"math/big"
	"time"
)

type Company struct {
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   int    `json:"phone"`
}

type CompanyCrudInterface interface {
	GetKey(id string) (Company, error)
	SetKey(value Company, id string, expiration time.Duration) error
	DelKey(key string) error
	GetAllValues() map[string]string
}

func (u *Company) GenerateId(compDomain Company) string {
	compByte, err := json.Marshal(compDomain)
	if err != nil {
		log.Println("Error on Marshalling on Generate id")
		return ""
	}
	urlHashBytes := sha256Of(string(compByte))
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Printf("Error during encoding %s", err)
		return ""
	}
	return string(encoded)
}
