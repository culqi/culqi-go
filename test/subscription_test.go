package culqi_test

import (
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestSubscription_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
	  "card_id": "crd_live_b3MMECR8cJ5tZqf2",
	  "plan_id": "pln_live_JucycAfngGozsHvC"
	}`)

	res, err := culqi.CreateSubscription(jsonData)
	if err != nil {
		t.Fatalf("Subscription.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}
}

func TestSubscription_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	res, err := culqi.GetByIDSubscription("sub_test_QjfgsHMPghROVZXa")
	if err != nil {
		t.Fatalf("Subscription.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}
}

func TestSubscription_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	params := url.Values{}
	params.Set("limit", "4")

	res, err := culqi.GetAllSubscription(params)
	if err != nil {
		t.Fatalf("Subscription.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscriptionAll = nil; want non-nil value")
	}
}

func TestSubscription_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
		"metadata": {
		"cliente_id": "259",
		"documento_identidad": "000551337"
		}
	}`)
	res, err := culqi.UpdateSubscription("sub_test_QjfgsHMPghROVZXa", jsonData)
	if err != nil {
		t.Fatalf("Subscription.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}
}

func TestSubscription_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	err := culqi.DeleteSubscriptions("sub_test_eY5TIGm70OMiCW89")
	if err != nil {
		t.Fatalf("Subscription.Delete() err = %v; want = %v", err, nil)
	}
}
