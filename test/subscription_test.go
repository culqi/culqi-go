package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
	util "github.com/culqi/culqi-go/utils"
)

func GetJsonSuscripcion() []byte {
	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)

	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)

	var jsonData = []byte(`{
		"plan_id": "` + idPlan + `",
		"card_id": "` + idCard + `"
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

func TestSubscription_Create(t *testing.T) {
	jsonData = GetJsonSuscripcion()

	_, res, err := culqi.CreateSubscription(jsonData)
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
	_, res, err := culqi.GetByIDSubscription(idSuscripcion, jsonData)
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

	_, res, err := culqi.GetAllSubscription(params, jsonData)
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
		"metadata": {
		"cliente_id": "259",
		"documento_identidad": "000551337"
		}
	}`)
	_, res, err := culqi.UpdateSubscription(idSuscripcion, jsonData)
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
	_, res, err := culqi.DeleteSubscriptions(idSuscripcion, jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Subscription.Delete() err = %v; want = %v", err, nil)
	}
}
