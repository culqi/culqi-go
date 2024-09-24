package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCard_Create(t *testing.T) {
	var json []byte
	json = GetJsonCard()
	fmt.Println(json)

	_, res, err := culqi.CreateCard(json)
	if err != nil {
		t.Fatalf("Card.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_CreateEncrypt(t *testing.T) {
	var json []byte
	json = GetJsonCard(encryptiondData...)
	fmt.Println(json)

	_, res, err := culqi.CreateCard(json, encryptiondData...)
	if err != nil {
		t.Fatalf("Card.CreateEncrypt() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_GetByID(t *testing.T) {
	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)

	var jsonData = []byte(``)
	_, res, err := culqi.GetByIDCard(idCard, jsonData)
	if err != nil {
		t.Fatalf("Card.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_GetAll(t *testing.T) {
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("limit", "4")

	_, res, err := culqi.GetAllCard(params, jsonData)
	if err != nil {
		t.Fatalf("Card.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCardAll = nil; want non-nil value")
	}
}

func TestCard_Update(t *testing.T) {
	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)
	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	_, res, err := culqi.UpdateCard(idCard, jsonData)
	if err != nil {
		t.Fatalf("Card.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_UpdateEncrypt(t *testing.T) {
	var idCard string
	idCard = GetIdCard(encryptiondData...)
	fmt.Println(idCard)
	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	_, res, err := culqi.UpdateCard(idCard, jsonData, encryptiondData...)
	if err != nil {
		t.Fatalf("Card.UpdateEncrypt() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_Delete(t *testing.T) {
	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)
	var jsonData = []byte(``)
	_, _, err := culqi.DeleteCard(idCard, jsonData)
	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}
}
