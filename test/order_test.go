package culqi_test

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"

	culqi "github.com/culqi/culqi-go"
	util "github.com/culqi/culqi-go/utils"
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
		"email": "prueba@gmail.com",
		"phone_number": "+51945145280"
	  },
	  "expiration_date":` + strconv.FormatInt(twoDaysLater.Unix(), 10) + `,
	  "confirm": "false"
	}`)
	return jsonData
}

func GetIdOrder() string {
	jsonData = getJsonData()

	_, res1, _ := culqi.CreateOrder(jsonData)

	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}

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
