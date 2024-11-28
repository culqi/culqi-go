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

	code, res, err := culqi.CreateCard(json)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Card.Create() err = %v; want = %v", err, nil)
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
	code, res, err := culqi.GetByIDCard(idCard, jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
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

	code, res, err := culqi.GetAllCard(params, jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
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
	code, res, err := culqi.UpdateCard(idCard, jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Card.Update() err = %v; want = %v", err, nil)
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
	code, _, err := culqi.DeleteCard(idCard, jsonData)
	fmt.Println(code)
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}
}
