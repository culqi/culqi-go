package culqi_test

import (
	"net/url"
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
		"source_id":     "tkn_test_WIouDPBhQH9OcKE8",
		"description":   "Curso GO desde Cero"		
	}`)

	res, err := culqi.CreateCharge(jsonData)
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
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
}
