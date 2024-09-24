package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCustomer_Create(t *testing.T) {
	var json []byte
	json = GetJsonCustomer()
	fmt.Println(json)
	_, res, err := culqi.CreateCustomer(json)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Customer.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}
}

func TestCustomer_CreateEncrypt(t *testing.T) {
	var json []byte
	json = GetJsonCustomer()
	fmt.Println(json)
	_, res, err := culqi.CreateCustomer(json, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Customer.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}
}

func TestCustomer_GetByID(t *testing.T) {
	var idCustomer string
	idCustomer = GetIdCustomer()
	fmt.Println(idCustomer)

	var jsonData = []byte(``)
	_, res, err := culqi.GetByIDCustomer(idCustomer, jsonData)
	if err != nil {
		t.Fatalf("Customer.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}
}

func TestCustomer_GetAll(t *testing.T) {
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("limit", "4")

	_, res, err := culqi.GetAllCustomer(params, jsonData)
	if err != nil {
		t.Fatalf("Customer.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomerAll = nil; want non-nil value")
	}
}

func TestCustomer_Update(t *testing.T) {
	var idCustomer string
	idCustomer = GetIdCustomer()
	fmt.Println(idCustomer)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	_, res, err := culqi.UpdateCustomer(idCustomer, jsonData)
	if err != nil {
		t.Fatalf("Customer.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}
}

func TestCustomer_UpdateEncrypt(t *testing.T) {
	var idCustomer string
	idCustomer = GetIdCustomer(encryptiondData...)
	fmt.Println(idCustomer)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	_, res, err := culqi.UpdateCustomer(idCustomer, jsonData, encryptiondData...)
	if err != nil {
		t.Fatalf("Customer.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}
}

func TestCustomer_Delete(t *testing.T) {
	var idCustomer string
	idCustomer = GetIdCustomer()
	fmt.Println(idCustomer)

	var jsonData = []byte(``)
	_, _, err := culqi.DeleteCustomer(idCustomer, jsonData)
	if err != nil {
		t.Fatalf("Customer.Delete() err = %v; want = %v", err, nil)
	}
}
