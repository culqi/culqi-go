package culqi_test

import (
	culqi "github.com/culqi/culqi-go"
	"net/url"
	"testing"
)

func TestOrder_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	var jsonData = []byte(`{
	  "amount": 12000,
	  "currency_code": "PEN", 
	  "description": "Venta de prueba", 
	  "order_number": "pedido299444343543344992",
	  "client_details": {
		"first_name": "Alexis", 
		"last_name": "Pumayalla", 
		"email": "apumayallag@gmail.com", 
		"phone_number": "+51945145280"
	  }, 
	  "expiration_date": 1673186377,
	  "confirm": "false"
	}`)
	res, err := culqi.CreateOrder(jsonData)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	res, err := culqi.GetByIDOrder("ord_test_HdkYBoii9Re5AOam")
	if err != nil {
		t.Fatalf("Order.GetByID() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	params := url.Values{}
	params.Set("paid", "false")

	res, err := culqi.GetAllOrder(params)
	if err != nil {
		t.Fatalf("Order.GetAll() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}

func TestOrder_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
		"expiration_date": 1661117022,
		"metadata": {
		"dni": "71701978"
		}
	}`)
	res, err := culqi.UpdateOrder("ord_test_4gLWlnFkNQB4iYhB", jsonData)
	if err != nil {
		t.Fatalf("Order.Update() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
func TestOrder_Confirm(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	res, err := culqi.ConfirmOrder("ord_test_HdkYBoii9Re5AOam")
	if err != nil {
		t.Fatalf("Order.Confirm() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}

func TestOrder_ConfirmTipo(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey)

	//create array
	var jsonData = []byte(`{
		"id": "ord_test_xjmEW4dIyJM9G4cc",
		"order_types": [
		"cuotealo",
		"cip"
		]
	}`)
	res, err := culqi.ConfirmTipoOrder(jsonData)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	res, err := culqi.DeleteOrder("ord_test_MrcA99oLfIRP0fKP")
	if err != nil {
		t.Fatalf("Order.Delete() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
