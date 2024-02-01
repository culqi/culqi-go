package culqi_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	culqi "github.com/culqi/culqi-go"
	util "github.com/culqi/culqi-go/utils"
)

// token
var jsonData = []byte(`{
	"card_number": "4111111111111111",
	"cvv": "123",
	"expiration_month": "09",
	"expiration_year": "2025",
	"email": "prueba1` + strconv.FormatInt(time.Now().UnixNano()/1000000, 10) + `@culqi.com"
}`)

// token
var jsonDataSubscription = []byte(`{
	"card_id": "crd_live_************",
	"plan_id": "pln_live_************",
	"tyc": true,
	"metadata":{
		"reference": "123123123"
	}
}`)

var jsonDataYape = []byte(`{		
	"amount": "700",
	"number_phone": "900000001",
	"otp": "425251"
}`)

func GetIdToken() string {
	_, res1, _ := culqi.CreateToken(jsonData)
	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}

// cargo
func GetJsonCharge(id string) []byte {
	msec := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	mapDataCargo := map[string]interface{}{
		"amount":        300,
		"capture":       true,
		"currency_code": "PEN",
		"email":         "test" + msec + "@aj.rdrgz",
		"source_id":     id,
		"description":   "Curso GO desde Cero",
	}
	jsonStr, _ := json.Marshal(mapDataCargo)
	return jsonStr
}

func GetIdCharge() string {
	var idToken string
	idToken = GetIdToken()

	var json []byte
	json = GetJsonCharge(idToken)

	_, res1, _ := culqi.CreateCharge(json)
	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}

// order
func getJsonData() (json []byte) {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(1000000000)
	//now := time.Now()
	//twoDaysLater := now.Add(time.Hour * 24 * 2)
	var jsonData = []byte(`{
	  "amount": 12000,
	  "currency_code": "PEN",
	  "description": "Venta de prueba",
	  "order_number": "pedido` + strconv.Itoa(number) + `",
	  "client_details": {
		"first_name": "Demo",
		"last_name": "Demo",
		"email": "prueba` + strconv.Itoa(number) + `@gmail.com",
		"phone_number": "+51945145280"
	  },
	  "expiration_date": "1893474000",
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

// customer
func GetJsonCustomer() []byte {
	msec := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	mapDataCustomer := map[string]interface{}{
		"first_name":   "Ejemplo",
		"last_name":    "Prueba",
		"email":        "prueba" + msec + "@culqi.com",
		"address":      "direccion",
		"address_city": "ciudad",
		"country_code": "PE",
		"phone_number": "987345123",
	}
	jsonStr, _ := json.Marshal(mapDataCustomer)
	return jsonStr
}

func GetIdCustomer() string {

	var json []byte
	json = GetJsonCustomer()

	_, res1, _ := culqi.CreateCustomer(json)

	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}

// card
func GetJsonCard() []byte {
	var idToken string
	idToken = GetIdToken()

	var idCustomer string
	idCustomer = GetIdCustomer()
	fmt.Println(idCustomer)

	mapDataCustomer := map[string]interface{}{
		"customer_id": idCustomer,
		"token_id":    idToken,
	}
	jsonStr, _ := json.Marshal(mapDataCustomer)
	return jsonStr
}

func GetIdCard() string {
	var json []byte
	json = GetJsonCard()

	_, res1, _ := culqi.CreateCard(json)

	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}

// plan
var jsonDataPlan = []byte(`{
    "short_name": "cp-prueb2442",
    "description": "Cypress PCI | ERRROR NO USAR",
    "amount": 300,
    "currency": "PEN",
    "interval_unit_time": 1,
    "interval_count": 1,
    "initial_cycles": {
      "count": 1,
      "has_initial_charge": true,
      "amount": 400,
      "interval_unit_time": 1
    },
    "name": "CY PCI - ERROR 100015",
    "image": "https://recurrencia-suscripciones-qa.s3.amazonaws.com/f097e1d5-e365-42f3-bc40-a27beab80f54",
	"metadata":{
		"key": "value"
	}
}`)

var jsonDataUpdatePlan = []byte(`{
    "short_name": "cp-prueb2442",
    "description": "Cypress PCI | ERRROR NO USAR",
    "status": 2,
    "name": "CY PCI - ERROR 100012",
	"metadata":{
		"key": "value"
	}
}`)

func GetIdPlan() string {
	_, res1, _ := culqi.CreatePlan(jsonDataPlan)
	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}

// Suscripción
func GetJsonSuscripcion() []byte {

	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)

	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)

	jsonData := []byte(`{
		"cardToken": "` + idCard + `",
		"planId": "` + idPlan + `",
		"customer": {},
		"metadata": {
			"envTest": "Autogenerado de Cypress"
		},
		"merchantId": "pk_live_0c301d16d8b892db",
		"tyc": true,
		"isPublicApi": true
	}`)

	return jsonData
}

func GetIdSuscripcion() string {
	jsonData = GetJsonSuscripcion()

	_, res1, _ := culqi.CreateSubscription(jsonData)

	var mapData map[string]interface{}
	mapData = util.JsonToMap([]byte(res1))
	id := fmt.Sprintf("%v", mapData["id"])

	return id
}
