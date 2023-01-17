package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCard_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	var jsonData = []byte(`{		
		"customer_id": "cus_test_AFIKH1gq8w7W7d4Q",
		"token_id":    "tkn_test_DY0oW1Edpu8vZWzM"
	}`)

	res, err := culqi.CreateCard(jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Card.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	res, err := culqi.GetByIDCard("crd_test_DFYNuxRluskZcbPZ", jsonData)
	if err != nil {
		t.Fatalf("Card.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("limit", "4")

	res, err := culqi.GetAllCard(params, jsonData)
	if err != nil {
		t.Fatalf("Card.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCardAll = nil; want non-nil value")
	}
}

func TestCard_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}
	culqi.Key(secretKey)
	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	res, err := culqi.UpdateCard("crd_test_DFYNuxRluskZcbPZ", jsonData)
	if err != nil {
		t.Fatalf("Card.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	err := culqi.DeleteCard("crd_test_DFYNuxRluskZcbPZ", jsonData)
	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}
}
