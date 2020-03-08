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

	culqi.Key(publicKey, secretKey)
	c := culqi.Card{
		CustomerID: "cus_test_XBpeiZRN49fZRofA",
		TokenID:    "tkn_test_m5YOT23kaGf8vCQy",
		Validate:   true,
		Metadata:   map[string]string{"pais": "Colombia"},
	}

	res, err := c.Create()
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

	culqi.Key(publicKey, secretKey)

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

	culqi.Key(publicKey, secretKey)

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

	culqi.Key(publicKey, secretKey)

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

	culqi.Key(publicKey, secretKey)

	p := culqi.Card{}
	err := p.Delete("crd_test_Qe0HG7VTfmTdFvgr")
	if err != nil {
		t.Fatalf("Card.Delete() err = %v; want = %v", err, nil)
	}
}
