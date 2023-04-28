package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

var jsonDataCargo = []byte(`{		
	"amount":      36200,
	"capture": true,
	"currency_code": "PEN",
	"email":         "test@aj.rdrgz",
	"source_id":     "tkn_test_XHqSZQtniMR547PE",
	"description":   "Curso GO desde Cero"		
}`)

func TestCharge_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	res, err := culqi.CreateCharge(jsonDataCargo)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

}

func TestCharge_CreateEncrypt(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	res, err := culqi.CreateCharge(jsonDataCargo, encryptiondData...)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

}

func TestCharge_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	res, err := culqi.GetByICharge("chr_test_zNHmuWgjkhpT0ke0", jsonData)
	if err != nil {
		t.Fatalf("Charge.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}
}

func TestCharge_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(``)
	params := url.Values{}
	params.Set("paid", "false")

	res, err := culqi.GetAllCharge(params, jsonData)
	if err != nil {
		t.Fatalf("Charge.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
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
	res, err := culqi.UpdateCharge("chr_test_zNHmuWgjkhpT0ke0", jsonData)
	if err != nil {
		t.Fatalf("Charge.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponseChargeAll = nil; want non-nil value")
	}
}
