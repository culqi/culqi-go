package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
	util "github.com/culqi/culqi-go/utils"
)

var jsonData = []byte(`{		
	"card_number": "4111111111111111",
	"cvv": "123",
	"expiration_month": "09",
	"expiration_year": "2025",
	"email": "prueba1@culqi.com"
}`)

var jsonDataYape = []byte(`{		
	"amount":      700,
	"number_phone": "900000001",
	"otp":         "425251"
}`)

func GetIdToken() string {
	_, res1, _ := culqi.CreateToken(jsonData)
	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}

func TestToken_Create(t *testing.T) {
	_, res, err := culqi.CreateToken(jsonData)
	fmt.Println(res)

	if err != nil {
		t.Fatalf("Token.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseToken = nil; want non-nil value")
	}
}

func TestToken_CreateEncrypt(t *testing.T) {
	_, res, err := culqi.CreateToken(jsonData, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Token.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseToken = nil; want non-nil value")
	}
}

func TestToken_CreateYape(t *testing.T) {
	_, res, err := culqi.CreateYape(jsonDataYape)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Token.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseTokenYape = nil; want non-nil value")
	}

}
func TestToken_Update(t *testing.T) {
	var id string
	id = GetIdToken()

	var jsonData = []byte(`{
	  "metadata": {
		 "dni": "4312354"
	   }
	}`)
	_, res, err := culqi.UpdateToken(id, jsonData)
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
	var id string
	id = GetIdToken()

	var jsonData = []byte(``)
	_, res, err := culqi.GetByIDToken(id, jsonData)
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
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("paid", "false")

	_, res, err := culqi.GetAllToken(params, jsonData)
	if err != nil {
		t.Fatalf("Order.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
