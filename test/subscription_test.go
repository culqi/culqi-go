package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestSubscription_Create(t *testing.T) {
	var funcName string = "TestSubscription_Create"
	logStartTest(funcName)

	jsonData = GetJsonSuscripcion()
	code, res, err := culqi.CreateSubscription(jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Subscription.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestSubscription_CreateEncrypt(t *testing.T) {
	var funcName string = "TestSubscription_CreateEncrypt"
	logStartTest(funcName)

	jsonData = GetJsonSuscripcion(encryptiondData...)
	code, res, err := culqi.CreateSubscription(jsonData, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Subscription.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestSubscription_GetByID(t *testing.T) {
	var funcName string = "TestSubscription_GetByID"
	logStartTest(funcName)

	var idSuscripcion string
	idSuscripcion = GetIdSuscripcion()
	fmt.Println(idSuscripcion)

	var jsonData = []byte(``)
	code, res, err := culqi.GetByIDSubscription(idSuscripcion, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Subscription.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestSubscription_GetAll(t *testing.T) {
	var funcName string = "TestSubscription_GetAll"
	logStartTest(funcName)

	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("limit", "4")

	code, res, err := culqi.GetAllSubscription(params, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Subscription.GetAll() err = %v; want = %v", err, nil)
	}
	if res == "" {
		t.Fatalf("ResponseSubscriptionAll = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestSubscription_Update(t *testing.T) {
	var funcName string = "TestSubscription_Update"
	logStartTest(funcName)

	var idSuscripcion string
	idSuscripcion = GetIdSuscripcion()
	fmt.Println(idSuscripcion)

	var jsonData = []byte(`{
		"card_id": "crd_live_***********",
		"metadata": {
			"cliente_id": "259",
			"documento_identidad": "000551337"
		}
	}`)
	code, res, err := culqi.UpdateSubscription(idSuscripcion, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Subscription.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestSubscription_UpdateEncrypt(t *testing.T) {
	var funcName = "TestSubscription_UpdateEncrypt"
	logStartTest(funcName)

	var idCard string
	idCard = GetIdCard(encryptiondData...)
	fmt.Println(idCard)

	var idSuscripcion string
	idSuscripcion = GetIdSuscripcion(encryptiondData...)
	fmt.Println(idSuscripcion)

	var jsonData = []byte(`{
		"card_id": "` + idCard + `",
		"metadata": {
			"cliente_id": "259",
			"documento_identidad": "000551337"
		}
	}`)

	code, res, err := culqi.UpdateSubscription(idSuscripcion, jsonData, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Subscription.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestSubscription_Delete(t *testing.T) {
	var funcName = "TestSubscription_Delete"
	logStartTest(funcName)

	var idSuscripcion string
	idSuscripcion = GetIdSuscripcion()
	fmt.Println(idSuscripcion)

	var jsonData = []byte(``)
	code, res, err := culqi.DeleteSubscriptions(idSuscripcion, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Subscription.Delete() err = %v; want = %v", err, nil)
	}

	logEndTest(funcName)
}
