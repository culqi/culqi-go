package culqi_test

import (
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
	c := culqi.Charge{
		Amount:       10100, // Monto del cargo. Sin punto decimal Ejemplo: 100.00 serían 10000
		Capture:      true,
		CurrencyCode: "PEN",
		Email:        "test@aj.rdrgz",
		SourceID:     "tkn_test_WIouDPBhQH9OcKE8",
		Description:  "Curso GO desde Cero",
		Metadata:     map[string]string{"user_id": "777"},
	}

	res, err := c.Create()
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	if res.Outcome.Type != "venta_exitosa" {
		t.Errorf("Charge.Outcome.Type = %s; want = %q", res.Outcome.Type, "venta_exitosa")
	}

	if !strings.HasPrefix(res.ID, "chr_test_") {
		t.Errorf("Charge.ID = %s; want prefix = %q", res.ID, "chr_test_")
	}
}

func TestCharge_Create3DS(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	c := culqi.Charge{
		Amount:       10100, // Monto del cargo. Sin punto decimal Ejemplo: 100.00 serían 10000
		Capture:      true,
		CurrencyCode: "PEN",
		Email:        "test@aj.rdrgz",
		SourceID:     "tkn_test_WIouDPBhQH9OcKE8",
		Description:  "Curso GO desde Cero",
		Metadata:     map[string]string{"user_id": "777"},
	}

	res, err := c.Create()
	if err != nil {
		t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseCharge = nil; want non-nil value")
	}

	if res.Outcome.Type != "venta_exitosa" {
		t.Errorf("Charge.Outcome.Type = %s; want = %q", res.Outcome.Type, "venta_exitosa")
	}

	if strings.HasPrefix(res.ReviewCode, "REVIEW") {

		d := culqi.Charge{
			Amount:            10100, // Monto del cargo. Sin punto decimal Ejemplo: 100.00 serían 10000
			Capture:           true,
			CurrencyCode:      "PEN",
			Email:             "test@aj.rdrgz",
			SourceID:          "tkn_test_WIouDPBhQH9OcKE8",
			Description:       "Curso GO desde Cero",
			Authentication3DS: map[string]string{"xid": "MTIzNDU2Nzg5MDEyMzQ1Njc40TA=", "cavv": "MTIzNDU2Nzg5MDEyMzQ1Njc40TA=", "directoryServerTransactionId": "5a636655-039f-4046-9564-50c084e6da85", "eci": "05", "protocolVersion": "2.2.0"},
		}

		res2, err := d.Create()
		if err != nil {
			t.Fatalf("Charge.Create() err = %v; want = %v", err, nil)
		}
		if res2 == nil {
			t.Fatalf("ResponseCard = nil; want non-nil value")
		}
	}
}

func TestCharge_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)

	c := culqi.Charge{}
	res, err := c.GetByID("chr_test_XCpvMiPnjQQpTIIv")
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

	c := culqi.Charge{}
	params := url.Values{}
	params.Set("paid", "false")

	res, err := c.GetAll(params)
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

	c := culqi.Charge{}
	res, err := c.Update("chr_test_XCpvMiPnjQQpTIIv", map[string]string{"orden_id": "789"})
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
