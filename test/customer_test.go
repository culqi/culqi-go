package culqi_test

import (
	"net/url"
	"strings"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCustomer_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
	  "first_name": "Ejemplo",
	  "last_name": "Prueba",
	  "email": "soportwwe1@culqi.com",
	  "address": "direccion",
	  "address_city": "ciudad",
	  "country_code": "PE",
	  "phone_number": "987345123"
	}`)

	res, err := culqi.CreateCustomer(jsonData)
	if err != nil {
		t.Fatalf("Customer.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "cus_test_") {
		t.Errorf("Customer.ID = %s; want prefix = %q", res.ID, "cus_test_")
	}
}

func TestCustomer_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	res, err := culqi.GetByIDCustomer("cus_test_XBpeiZRN49fZRofA")
	if err != nil {
		t.Fatalf("Customer.GetByID() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}
}

func TestCustomer_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	params := url.Values{}
	params.Set("limit", "4")

	res, err := culqi.GetAllCustomer(params)
	if err != nil {
		t.Fatalf("Customer.GetAll() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCustomerAll = nil; want non-nil value")
	}
}

func TestCustomer_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	res, err := culqi.UpdateCustomer("cus_test_XBpeiZRN49fZRofA", jsonData)
	if err != nil {
		t.Fatalf("Customer.Update() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCustomer = nil; want non-nil value")
	}

	if res.Metadata["nid"] != "1122334455" {
		t.Errorf(`Customer.Metadata["nid"] = %s; want = %q`, res.Metadata["nid"], "1122334455")
	}
}

func TestCustomer_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	err := culqi.DeleteCustomer("cus_test_wuCZZ9jqhO0RY6xZ")
	if err != nil {
		t.Fatalf("Customer.Delete() err = %v; want = %v", err, nil)
	}
}
