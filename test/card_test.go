package culqi_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
	util "github.com/culqi/culqi-go/utils"
)

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

func TestCard_Create(t *testing.T) {
	var json []byte
	json = GetJsonCard()
	fmt.Println(json)

	_, res, err := culqi.CreateCard(json)
	if err != nil {
		t.Fatalf("Card.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_GetByID(t *testing.T) {
	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)

	var jsonData = []byte(``)
	_, res, err := culqi.GetByIDCard(idCard, jsonData)
	if err != nil {
		t.Fatalf("Card.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_GetAll(t *testing.T) {
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("limit", "4")

	_, res, err := culqi.GetAllCard(params, jsonData)
	if err != nil {
		t.Fatalf("Card.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCardAll = nil; want non-nil value")
	}
}

func TestCard_Update(t *testing.T) {
	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)
	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	_, res, err := culqi.UpdateCard(idCard, jsonData)
	if err != nil {
		t.Fatalf("Card.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_Delete(t *testing.T) {
	var idCard string
	idCard = GetIdCard()
	fmt.Println(idCard)
	var jsonData = []byte(``)
	_, _, err := culqi.DeleteCard(idCard, jsonData)
	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}
}
