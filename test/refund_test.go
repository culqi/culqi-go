package culqi_test

import (
	"fmt"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestRefund_Create(t *testing.T) {
	var funcName string = "TestRefund_Create"
	logStartTest(funcName)

	var json []byte
	json = GetJsonRefund()
	fmt.Println(json)

	code, res, err := culqi.CreateRefund(json)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Refund.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestRefund_CreateEncrypt(t *testing.T) {
	var funcName string = "TestRefund_CreateEncrypt"
	logStartTest(funcName)

	var json []byte
	json = GetJsonRefund(encryptiondData...)
	fmt.Println(json)

	code, res, err := culqi.CreateRefund(json, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Refund.CreateEncrypt() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestRefund_Update(t *testing.T) {
	var funcName string = "TestRefund_Update"
	logStartTest(funcName)

	var idRefund string
	idRefund = GetIdRefund()
	fmt.Println(idRefund)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)

	code, res, err := culqi.UpdateRefund(idRefund, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Refund.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestRefund_UpdateEncrypt(t *testing.T) {
	var funcName string = "TestRefund_UpdateEncrypt"
	logStartTest(funcName)

	var idRefund string
	idRefund = GetIdRefund(encryptiondData...)
	fmt.Println(idRefund)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)

	code, res, err := culqi.UpdateRefund(idRefund, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Refund.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}

	logEndTest(funcName)
}
