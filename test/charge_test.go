package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCharge_Create(t *testing.T) {
	var funcName string = "TestCharge_Create"
	logStartTest(funcName)

	var idToken string
	idToken = GetIdToken()
	var json []byte
	json = GetJsonCharge(idToken)

	code, res, err := culqi.CreateCharge(json)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_RecurrentHeader(t *testing.T) {
	var funcName string = "TestCharge_RecurrentHeader"
	logStartTest(funcName)

	var idToken string
	idToken = GetIdToken()
	var json []byte
	var optiomalParams []byte
	json = GetJsonCharge(idToken)
	optiomalParams = []byte(`{
				"custom_headers": {
					"X-Charge-Channel": "recurrent"
				}
			}`)

	code, res, err := culqi.CreateCharge(json, optiomalParams...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_CreateEncrypt(t *testing.T) {
	var funcName string = "TestCharge_CreateEncrypt"
	logStartTest(funcName)

	var idToken string
	idToken = GetIdToken(encryptiondData...)
	var json []byte
	json = GetJsonCharge(idToken)

	code, res, err := culqi.CreateCharge(json, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_Capture(t *testing.T) {
	var funcName string = "TestCharge_Capture"
	logStartTest(funcName)

	var id string
	id = GetIdCharge()

	var jsonData = []byte(``)
	code, res, err := culqi.ChargeCapture(id, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.Capture() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_CaptureEncrypt(t *testing.T) {
	var funcName string = "TestCharge_CaptureEncrypt"
	logStartTest(funcName)

	var id string
	id = GetIdCharge(encryptiondData...)

	var jsonData = []byte(``)
	code, res, err := culqi.ChargeCapture(id, jsonData, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.Capture() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_GetByID(t *testing.T) {
	var funcName string = "TestCharge_GetByID"
	logStartTest(funcName)

	var id string
	id = GetIdCharge()

	var jsonData = []byte(``)
	code, res, err := culqi.GetByIdCharge(id, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_GetAll(t *testing.T) {
	var funcName string = "TestCharge_GetAll"
	logStartTest(funcName)

	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("paid", "false")

	code, res, err := culqi.GetAllCharge(params, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_Update(t *testing.T) {
	var funcName string = "TestCharge_Update"
	logStartTest(funcName)

	var id string
	id = GetIdCharge()

	var jsonData = []byte(`{
	"metadata": {
		"documentType": "1",
		"documentNumber": "99999999"
		}
	}`)
	code, res, err := culqi.UpdateCharge(id, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestCharge_UpdateEncrypt(t *testing.T) {
	var funcName string = "TestCharge_UpdateEncrypt"
	logStartTest(funcName)

	var id string
	id = GetIdCharge(encryptiondData...)

	var jsonData = []byte(`{
	"metadata": {
		"documentType": "1",
		"documentNumber": "99999999"
		}
	}`)
	code, res, err := culqi.UpdateCharge(id, jsonData, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Charge.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}

	logEndTest(funcName)
}
