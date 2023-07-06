package culqi_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
	util "github.com/culqi/culqi-go/utils"
)

/*
	var jsonDataCargo = []byte(`{
		"amount":      300,
		"capture": true,
		"currency_code": "PEN",
		"email":         "test@aj.rdrgz",
		"source_id":     "tkn_test_XHqSZQtniMR547PE",
		"description":   "Curso GO desde Cero"
	}`)
*/
func GetJsonCharge(id string) []byte {
	mapDataCargo := map[string]interface{}{
		"amount":        300,
		"capture":       true,
		"currency_code": "PEN",
		"email":         "test@aj.rdrgz",
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

func TestCharge_Create(t *testing.T) {
	var idToken string
	idToken = GetIdToken()
	var json []byte
	json = GetJsonCharge(idToken)

	_, res, err := culqi.CreateCharge(json)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}
}

func TestCharge_CreateEncrypt(t *testing.T) {
	var idToken string
	idToken = GetIdToken()
	var json []byte
	json = GetJsonCharge(idToken)

	_, res, err := culqi.CreateCharge(json, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}
}

func TestCharge_GetByID(t *testing.T) {
	var id string
	id = GetIdCharge()

	var jsonData = []byte(``)
	_, res, err := culqi.GetByIdCharge(id, jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Charge.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}
}

func TestCharge_GetAll(t *testing.T) {
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("paid", "false")

	_, res, err := culqi.GetAllCharge(params, jsonData)
	if err != nil {
		t.Fatalf("Charge.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}
}

func TestCharge_Update(t *testing.T) {
	var id string
	id = GetIdCharge()

	var jsonData = []byte(`{
	"metadata": {
		"documentType": "1",
		"documentNumber": "99999999"
		}
	}`)
	_, res, err := culqi.UpdateCharge(id, jsonData)
	if err != nil {
		t.Fatalf("Charge.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}
}
