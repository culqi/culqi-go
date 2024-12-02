package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestSubscription_Create(t *testing.T) {
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
}

func TestSubscription_CreateEncrypt(t *testing.T) {
	jsonData = GetJsonSuscripcion(encryptiondData...)
	_, res, err := culqi.CreateSubscription(jsonData, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Subscription.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}
}

func TestSubscription_GetByID(t *testing.T) {
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
}

func TestSubscription_GetAll(t *testing.T) {
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
}

func TestSubscription_Update(t *testing.T) {
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
}

func TestSubscription_UpdateEncrypt(t *testing.T) {
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
	_, res, err := culqi.UpdateSubscription(idSuscripcion, jsonData, encryptiondData...)
	if err != nil {
		t.Fatalf("Subscription.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}
}

func TestSubscription_Delete(t *testing.T) {
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
}
