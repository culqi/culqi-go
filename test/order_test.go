package culqi_test

import (
	"fmt"
	culqi "github.com/culqi/culqi-go"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestOrder_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	c := culqi.Order{
		Amount:         10100, // Monto del cargo. Sin punto decimal Ejemplo: 100.00 serían 10000
		CurrencyCode:   "PEN",
		Description:    "Curso GO desde Cero",
		OrderNumber:    fmt.Sprint("pedido-go-", time.Now().UnixNano()/1e6),
		ClientDetails:  map[string]string{"first_name": "Richard", "last_name": "Hendricks", "email": "richard@piedpiper.com", "phone_number": "+51945145280"},
		ExpirationDate: 1673129486,
	}
	fmt.Println(c)
	res, err := c.Create()
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "ord_test_") {
		t.Errorf("Order.ID = %s; want prefix = %q", res.ID, "ord_test_")
	}
}

func TestOrder_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	c := culqi.Order{}
	res, err := c.GetByID("ord_test_HdkYBoii9Re5AOam")
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

	c := culqi.Order{}
	params := url.Values{}
	params.Set("paid", "false")

	res, err := c.GetAll(params)
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

	c := culqi.Order{}
	res, err := c.Update("ord_test_4gLWlnFkNQB4iYhB", map[string]string{"orden_id": "789"})
	if err != nil {
		t.Fatalf("Order.Update() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}

	if res.Metadata["orden_id"] != "789" {
		t.Errorf(`Order.Metadata["orden_id"] = %s; want = %q`, res.Metadata["orden_id"], "789")
	}
}
func TestOrder_Confirm(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	c := culqi.Order{}
	res, err := c.Confirm("ord_test_HdkYBoii9Re5AOam", map[string]string{"orden_id": "789"})
	if err != nil {
		t.Fatalf("Order.Confirm() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}

	if res.Metadata["orden_id"] != "789" {
		t.Errorf(`Order.Metadata["orden_id"] = %s; want = %q`, res.Metadata["orden_id"], "789")
	}
}

func TestOrder_ConfirmTipo(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey)

	//create array
	var metodos = [2]string{"cuotealo", "cip"}

	c := culqi.OrderTipo{
		OrderId:    "ord_test_H28xMwYpp0qTLWGt",
		OrderTypes: metodos,
	}
	res, err := c.ConfirmTipo()
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "ord_test_") {
		t.Errorf("Order.ID = %s; want prefix = %q", res.ID, "ord_test_")
	}

}

func TestOrder_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	c := culqi.Order{}
	res, err := c.Delete("ord_test_MrcA99oLfIRP0fKP", map[string]string{"orden_id": "789"})
	if err != nil {
		t.Fatalf("Order.Delete() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}

	if res.Metadata["orden_id"] != "789" {
		t.Errorf(`Order.Metadata["orden_id"] = %s; want = %q`, res.Metadata["orden_id"], "789")
	}
}
