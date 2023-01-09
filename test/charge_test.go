package culqi_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestCharge_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{		
		"amount":      36200,
		"capture": true,
		"currency_code": "PEN",
		"email":         "test@aj.rdrgz",
		"source_id":     "tkn_test_E0wpKOMJC4ljuNaw",
		"description":   "Curso GO desde Cero"		
	}`)

	res, err := culqi.CreateCharge(jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	if res.action_code != "REVIEW" {

		if res.Outcome.Type != "venta_exitosa" {
			t.Errorf("Charge.Outcome.Type = %s; want = %q", res.Outcome.Type, "venta_exitosa")
		}
	}

	if !strings.HasPrefix(res.ID, "chr_test_") {
		t.Errorf("Charge.ID = %s; want prefix = %q", res.ID, "chr_test_")
	}
}

func TestCharge_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	res, err := culqi.GetByICharge("chr_test_XCpvMiPnjQQpTIIv")
	if err != nil {
		t.Fatalf("Charge.GetByID() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}
}

func TestCharge_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	params := url.Values{}
	params.Set("paid", "false")

	res, err := culqi.GetAllCharge(params)
	if err != nil {
		t.Fatalf("Charge.GetAll() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}
}

func TestCharge_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
	"metadata": {
		"documentType": "1",
		"documentNumber": "99999999"
		}
	}`)
	res, err := culqi.UpdateCharge("chr_test_XCpvMiPnjQQpTIIv", jsonData)
	if err != nil {
		t.Fatalf("Charge.Update() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}

	if res.Metadata["orden_id"] != "789" {
		t.Errorf(`Charge.Metadata["orden_id"] = %s; want = %q`, res.Metadata["orden_id"], "789")
	}
}
