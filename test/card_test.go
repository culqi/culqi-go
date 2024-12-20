package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCard_Create(t *testing.T) {
	var funcName string = "TestCard_Create"
	logStartTest(funcName)

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

	logEndTest(funcName)
}

func TestCard_CreateEncrypt(t *testing.T) {
	var funcName string = "TestCard_CreateEncrypt"
	logStartTest(funcName)

	var json []byte
	json = GetJsonCard(encryptiondData...)
	fmt.Println(json)

	code, res, err := culqi.CreateCard(json, encryptiondData...)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Card.CreateEncrypt() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCard_GetByID(t *testing.T) {
	var funcName string = "TestCard_GetByID"
	logStartTest(funcName)

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
	logEndTest(funcName)
}

func TestCard_GetAll(t *testing.T) {
	var funcName string = "TestCard_GetAll"
	logStartTest(funcName)

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

	logEndTest(funcName)
}

func TestCard_Update(t *testing.T) {
	var funcName string = "TestCard_Update"
	logStartTest(funcName)

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

	logEndTest(funcName)
}

func TestCard_UpdateEncrypt(t *testing.T) {
	var funcName string = "TestCard_UpdateEncrypt"
	logStartTest(funcName)

	var idCard string
	idCard = GetIdCard(encryptiondData...)
	fmt.Println(idCard)
	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	code, res, err := culqi.UpdateCard(idCard, jsonData, encryptiondData...)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Card.UpdateEncrypt() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCard_Delete(t *testing.T) {
	var funcName string = "TestCard_Delete"
	logStartTest(funcName)

	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)
	var jsonData = []byte(``)
	code, res, err := culqi.DeleteCard(idCard, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}

	logEndTest(funcName)
}
