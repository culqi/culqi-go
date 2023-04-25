package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestToken_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey)
	var jsonData = []byte(`{		
		"card_number": "4111111111111111",
		"cvv": "123",
		"expiration_month": "07",
		"expiration_year": "2025",
		"email": "prueba1@culqi.com"
	}`)

	var encryptiondData = []byte(`{		
		"rsa_public_key": "-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYp0451xITpczkBrl5Goxkh7m1oynj8eDHypIn7HmbyoNJd8cS4OsT850hIDBwYmFuwmxF1YAJS8Cd2nes7fjCHh+7oNqgNKxM2P2NLaeo4Uz6n9Lu4KKSxTiIT7BHiSryC0+Dic91XLH7ZTzrfryxigsc+ZNndv0fQLOW2i6OhwIDAQAB
-----END PUBLIC KEY-----",
		"rsa_id": "508fc232-0a9d-4fc0-a192-364a0b782b89"
	}`)

	res, err := culqi.CreateToken(jsonData, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Token.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseToken = nil; want non-nil value")
	}
}

func TestToken_CreateYape(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey)
	var jsonData = []byte(`{		
		"amount":      36200,
		"fingerprint": "86d3c875769bf62b0471b47853bfda77",
		"number_phone": "900000001",
		"otp":         "425251"
	}`)

	res, err := culqi.CreateYape(jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Token.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseTokenYape = nil; want non-nil value")
	}

}
func TestToken_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
	  "metadata": {
		 "dni": "krthkrh"
	   }
	}`)
	res, err := culqi.UpdateToken("tkn_test_lgMNwCh5CBICTsGu", jsonData)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Token.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseTokenAll = nil; want non-nil value")
	}
}
func TestToken_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave secreta")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	res, err := culqi.GetByIDToken("tkn_test_lgMNwCh5CBICTsGu", jsonData)
	fmt.Println(err)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Order.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestToken_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("paid", "false")

	res, err := culqi.GetAllToken(params, jsonData)
	if err != nil {
		t.Fatalf("Order.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
