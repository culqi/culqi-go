package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestOrder_Create(t *testing.T) {
	jsonData = getJsonData()

	_, res, err := culqi.CreateOrder(jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_CreateEncrypt(t *testing.T) {
	jsonData = getJsonData()
	//encryptiondData = getEncryptionParams()

	_, res, err := culqi.CreateOrder(jsonData, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_GetByID(t *testing.T) {
	var idOrder string
	idOrder = GetIdOrder()
	fmt.Println(idOrder)

	var jsonData = []byte(``)
	_, res, err := culqi.GetByIDOrder(idOrder, jsonData)
	if err != nil {
		t.Fatalf("Order.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_GetAll(t *testing.T) {
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("paid", "false")

	_, res, err := culqi.GetAllOrder(params, jsonData)
	if err != nil {
		t.Fatalf("Order.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}

func TestOrder_Update(t *testing.T) {
	var idOrder string
	idOrder = GetIdOrder()
	fmt.Println(idOrder)

	var jsonData = []byte(`{
		"expiration_date": 1661117022,
		"metadata": {
		"dni": "71701978"
		}
	}`)
	_, res, err := culqi.UpdateOrder(idOrder, jsonData)
	if err != nil {
		t.Fatalf("Order.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
func TestOrder_Confirm(t *testing.T) {
	var idOrder string
	idOrder = GetIdOrder()
	fmt.Println(idOrder)

	var jsonData = []byte(``)
	_, res, err := culqi.ConfirmOrder(idOrder, jsonData)
	if err != nil {
		t.Fatalf("Order.Confirm() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}

func TestOrder_ConfirmTipo(t *testing.T) {
	var idOrder string
	idOrder = GetIdOrder()
	fmt.Println(idOrder)
	//create array
	var jsonData = []byte(`{
		"order_id": "` + idOrder + `",
		"order_types": [
		"cuotealo",
		"cip"
		]
	}`)
	_, res, err := culqi.ConfirmTipoOrder(jsonData)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_Delete(t *testing.T) {
	var idOrder string
	idOrder = GetIdOrder()
	fmt.Println(idOrder)

	var jsonData = []byte(``)
	_, _, err := culqi.DeleteOrder(idOrder, jsonData)
	if err != nil {
		t.Fatalf("Order.Delete() err = %v; want = %v", err, nil)
	}

}
