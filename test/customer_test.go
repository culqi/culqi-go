package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCustomer_Create(t *testing.T) {
	var funcName string = "TestCustomer_Create"
	logStartTest(funcName)

	var json []byte
	json = GetJsonCustomer()

	code, res, err := culqi.CreateCustomer(json)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Customer.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCustomer_CreateEncrypt(t *testing.T) {
	var funcName string = "TestCustomer_CreateEncrypt"
	logStartTest(funcName)

	var json []byte
	json = GetJsonCustomer()
	fmt.Println(json)
	code, res, err := culqi.CreateCustomer(json, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Customer.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCustomer_GetByID(t *testing.T) {
	var funcName string = "TestCustomer_GetByID"
	logStartTest(funcName)

	var idCustomer string
	idCustomer = GetIdCustomer()
	fmt.Println(idCustomer)

	var jsonData = []byte(``)
	code, res, err := culqi.GetByIDCustomer(idCustomer, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Customer.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCustomer_GetAll(t *testing.T) {
	var funcName string = "TestCustomer_GetAll"
	logStartTest(funcName)

	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("limit", "4")

	code, res, err := culqi.GetAllCustomer(params, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Customer.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomerAll = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCustomer_Update(t *testing.T) {
	var funcName string = "TestCustomer_Update"
	logStartTest(funcName)

	var idCustomer string
	idCustomer = GetIdCustomer()
	fmt.Println(idCustomer)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	code, res, err := culqi.UpdateCustomer(idCustomer, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Customer.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}
	logEndTest(funcName)
}

func TestCustomer_UpdateEncrypt(t *testing.T) {
	var funcName string = "TestCustomer_UpdateEncrypt"
	logStartTest(funcName)

	var idCustomer string
	idCustomer = GetIdCustomer(encryptiondData...)
	fmt.Println(idCustomer)

	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	code, res, err := culqi.UpdateCustomer(idCustomer, jsonData, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Customer.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCustomer_Delete(t *testing.T) {
	var funcName string = "TestCustomer_Delete"
	logStartTest(funcName)

	var idCustomer string
	idCustomer = GetIdCustomer()
	fmt.Println(idCustomer)

	var jsonData = []byte(``)
	code, res, err := culqi.DeleteCustomer(idCustomer, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Customer.Delete() err = %v; want = %v", err, nil)
	}

	logEndTest(funcName)
}
