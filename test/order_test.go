package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestOrder_Create(t *testing.T) {
	jsonData = getJsonData()

	code, res, err := culqi.CreateOrder(jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

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

	code, res, err := culqi.CreateOrder(jsonData, encryptiondData...)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
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
	code, res, err := culqi.GetByIDOrder(idOrder, jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
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

	code, res, err := culqi.GetAllOrder(params, jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
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
	expiration_date := generateTimeStamp()
	fmt.Println(idOrder)

	var jsonData = []byte(`{
		"expiration_date": "` + expiration_date + `",
		"metadata": {
		"dni": "77777777"
		}
	}`)
	code, res, err := culqi.UpdateOrder(idOrder, jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Order.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
func TestOrder_UpdateEncrypt(t *testing.T) {
	var idOrder string
	idOrder = GetIdOrder(encryptiondData...)
	fmt.Println(idOrder)

	var jsonData = []byte(`{
		"expiration_date": 1661117022,
		"metadata": {
		"dni": "71701978"
		}
	}`)
	_, res, err := culqi.UpdateOrder(idOrder, jsonData, encryptiondData...)
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
	code, res, err := culqi.ConfirmOrder(idOrder, jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
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
	code, res, err := culqi.ConfirmTipoOrder(jsonData)
	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)
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
	code, _, err := culqi.DeleteOrder(idOrder, jsonData)
	fmt.Println(code)
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Order.Delete() err = %v; want = %v", err, nil)
	}

}
