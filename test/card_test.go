package culqi_test

import (
	"fmt"
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
	c := culqi.Card{
		CustomerID: "cus_test_AFIKH1gq8w7W7d4Q",
		TokenID:    "tkn_test_RX7Hv23z9I8VEmZl",
	}

	res, err := c.Create()
	fmt.Println(err)
	fmt.Println(res)
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
func TestCard_Create3ds(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	c := culqi.Card{
		CustomerID: "cus_test_NtUGmY9Oyjr8oNX7",
		TokenID:    "tkn_test_Qr4iXvqZaqmahC9P",
	}

	res, err := c.Create()
	if err != nil {
		t.Fatalf("Card.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCard = nil; want non-nil value")
	}
	fmt.Println(res)
	if strings.HasPrefix(res.ReviewCode, "REVIEW") {

		d := culqi.Card{
			CustomerID:        "cus_test_NtUGmY9Oyjr8oNX7",
			TokenID:           "tkn_test_Qr4iXvqZaqmahC9P",
			Authentication3DS: map[string]string{"xid": "MTIzNDU2Nzg5MDEyMzQ1Njc40TA=", "cavv": "MTIzNDU2Nzg5MDEyMzQ1Njc40TA=", "directoryServerTransactionId": "5a636655-039f-4046-9564-50c084e6da85", "eci": "05", "protocolVersion": "2.2.0"},
		}

		res2, err := d.Create()
		if err != nil {
			t.Fatalf("Card.Create() err = %v; want = %v", err, nil)
		}
		if res2 == nil {
			t.Fatalf("ResponseCard = nil; want non-nil value")
		}
	}
}

func TestCard_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	c := culqi.Card{}
	res, err := c.GetByID("crd_test_Qe0HG7VTfmTdFvgr")
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

	c := culqi.Card{}
	params := url.Values{}
	params.Set("limit", "4")

	res, err := c.GetAll(params)
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

	c := culqi.Card{}
	res, err := c.Update("crd_test_Qe0HG7VTfmTdFvgr", map[string]string{"banco": "Bancolombia"})
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

	p := culqi.Card{}
	err := p.Delete("crd_test_Qe0HG7VTfmTdFvgr")
	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}
}
