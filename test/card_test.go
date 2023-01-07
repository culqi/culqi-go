package culqi_test

import (
	"net/url"
	"strings"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCard_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	var jsonData = []byte(`{		
		"customer_id": "cus_test_AFIKH1gq8w7W7d4Q",
		"token_id":    "tkn_test_RX7Hv23z9I8VEmZl"
	}`)

	res, err := culqi.CreateCard(jsonData)
	if err != nil {
		t.Fatalf("Card.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "crd_test_") {
		t.Errorf("Card.ID = %s; want prefix = %q", res.ID, "crd_test_")
	}
}

func TestCard_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	res, err := culqi.GetByIDCard("crd_test_Qe0HG7VTfmTdFvgr")
	if err != nil {
		t.Fatalf("Card.GetByID() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
}

func TestCard_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	params := url.Values{}
	params.Set("limit", "4")

	res, err := culqi.GetAllCard(params)
	if err != nil {
		t.Fatalf("Card.GetAll() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCardAll = nil; want non-nil value")
	}
}

func TestCard_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}
	culqi.Key(secretKey)
	var jsonData = []byte(`{
		"metadata": {
		"dni": "71702323"
		}
	}`)
	res, err := culqi.UpdateCard("crd_test_Qe0HG7VTfmTdFvgr", jsonData)
	if err != nil {
		t.Fatalf("Card.Update() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}

	if res.Metadata["banco"] != "Bancolombia" {
		t.Errorf(`Card.Metadata["banco"] = %s; want = %q`, res.Metadata["banco"], "Bancolombia")
	}
}

func TestCard_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	err := culqi.DeleteCard("crd_test_Qe0HG7VTfmTdFvgr")
	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}
}
