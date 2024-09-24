package culqi_test

import (
	"fmt"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestRefund_Create(t *testing.T) {
	var json []byte
	json = GetJsonRefund()
	fmt.Println(json)

	_, res, err := culqi.CreateRefund(json)
	if err != nil {
		t.Fatalf("Refund.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}
}

func TestRefund_CreateEncrypt(t *testing.T) {
	var json []byte
	json = GetJsonRefund(encryptiondData...)
	fmt.Println(json)

	_, res, err := culqi.CreateRefund(json, encryptiondData...)
	if err != nil {
		t.Fatalf("Refund.CreateEncrypt() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}
}

func TestRefund_Update(t *testing.T) {
	var idRefund string
	idRefund = GetIdRefund()
	fmt.Println(idRefund)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	_, res, err := culqi.UpdateRefund(idRefund, jsonData)
	if err != nil {
		t.Fatalf("Refund.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}
}

func TestRefund_UpdateEncrypt(t *testing.T) {
	var idRefund string
	idRefund = GetIdRefund(encryptiondData...)
	fmt.Println(idRefund)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	_, res, err := culqi.UpdateRefund(idRefund, jsonData)
	if err != nil {
		t.Fatalf("Refund.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseRefund = nil; want non-nil value")
	}
}
