package culqi_test

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"

	culqi "github.com/culqi/culqi-go"
)

func getJsonData() (json []byte) {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(1000000000)
	now := time.Now()
	twoDaysLater := now.Add(time.Hour * 24 * 2)
	var jsonData = []byte(`{
	  "amount": 12000,
	  "currency_code": "PEN", 
	  "description": "Venta de prueba", 
	  "order_number": "pedido` + strconv.Itoa(number) + `",
	  "client_details": {
		"first_name": "Alexis", 
		"last_name": "Pumayalla", 
		"email": "apumayallag@gmail.com", 
		"phone_number": "+51945145280"
	  }, 
	  "expiration_date":` + strconv.FormatInt(twoDaysLater.Unix(), 10) + `,
	  "confirm": "false"
	}`)
	return jsonData
}

func TestOrder_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	jsonData = getJsonData()
	fmt.Println(string(jsonData))
	res, err := culqi.CreateOrder(jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_CreateEncrypt(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	jsonData = getJsonData()
	//encryptiondData = getEncryptionParams()

	fmt.Println(string(jsonData))
	res, err := culqi.CreateOrder(jsonData, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	res, err := culqi.GetByIDOrder("ord_test_ozYVmT8qkAvBo8rg", jsonData)
	if err != nil {
		t.Fatalf("Order.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("paid", "false")

	res, err := culqi.GetAllOrder(params, jsonData)
	if err != nil {
		t.Fatalf("Order.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
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
	res, err := culqi.UpdateOrder("ord_test_ozYVmT8qkAvBo8rg", jsonData)
	if err != nil {
		t.Fatalf("Order.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}
func TestOrder_Confirm(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	res, err := culqi.ConfirmOrder("ord_test_hmiCojFBInTpbkEh", jsonData)
	if err != nil {
		t.Fatalf("Order.Confirm() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrderAll = nil; want non-nil value")
	}
}

func TestOrder_ConfirmTipo(t *testing.T) {
	if publicKey == "" {
		t.Skip("No se indicó una llave publica")
	}

	culqi.Key(publicKey)

	//create array
	var jsonData = []byte(`{
		"order_id": "ord_test_ySo6qf0esBjkfM6R",
		"order_types": [
		"cuotealo",
		"cip"
		]
	}`)
	res, err := culqi.ConfirmTipoOrder(jsonData)
	if err != nil {
		t.Fatalf("Order.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseOrder = nil; want non-nil value")
	}
}

func TestOrder_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	_, err := culqi.DeleteOrder("ord_test_cVpkdDWsLfDimBt9", jsonData)
	if err != nil {
		t.Fatalf("Order.Delete() err = %v; want = %v", err, nil)
	}

}
