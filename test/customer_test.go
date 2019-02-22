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

	culqi.Key(publicKey, secretKey)
	c := culqi.Customer{
		FirstName:   "Alejandro",
		LastName:    "Rodriguez",
		Email:       "test@aj.rdrgz",
		Address:     "Bogotá, Colombia",
		AddressCity: "Bogotá",
		CountryCode: "CO",
		PhoneNumber: "3207777777",
		Metadata:    map[string]string{"nid": "123456789"},
	}

	res, err := c.Create()
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

	culqi.Key(publicKey, secretKey)

	c := culqi.Customer{}
	res, err := c.GetByID("cus_test_XBpeiZRN49fZRofA")
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

	culqi.Key(publicKey, secretKey)

	c := culqi.Customer{}
	params := url.Values{}
	params.Set("limit", "4")

	res, err := c.GetAll(params)
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

	culqi.Key(publicKey, secretKey)

	c := culqi.Customer{}
	res, err := c.Update("cus_test_XBpeiZRN49fZRofA", map[string]string{"nid": "1122334455"})
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

	culqi.Key(publicKey, secretKey)

	p := culqi.Customer{}
	err := p.Delete("cus_test_wuCZZ9jqhO0RY6xZ")
	if err != nil {
		t.Fatalf("Customer.Delete() err = %v; want = %v", err, nil)
	}
}
