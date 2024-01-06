package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

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
	params.Set("device_type", "mobile")

	_, res, err := culqi.GetAllToken(params, jsonData)
	if err != nil {
		t.Fatalf("Order.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
