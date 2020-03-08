package culqi_test

import (
	"net/url"
	"strings"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestPlan_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)
	p := culqi.Plan{
		Name:          "Suscripción Premium",
		Amount:        3000, // Monto del plan a cobrar recurrentemente. Sin punto decimal Ejemplo: 30.00 serían 3000
		CurrencyCode:  "USD",
		Interval:      "meses",
		IntervalCount: 1,
		Metadata:      map[string]string{"descripción": "Plan premium black friday"},
	}

	res, err := p.Create()
	if err != nil {
		t.Fatalf("Plan.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "pln_test_") {
		t.Errorf("Plan.ID = %s; want prefix = %q", res.ID, "pln_test_")
	}
}

func TestPlan_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	p := culqi.Plan{}
	res, err := p.GetByID("pln_test_oFvWoKSAZOAH1weu")
	if err != nil {
		t.Fatalf("Plan.GetByID() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
}

func TestPlan_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	p := culqi.Plan{}
	params := url.Values{}
	params.Set("limit", "4")

	res, err := p.GetAll(params)
	if err != nil {
		t.Fatalf("Plan.GetAll() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponsePlanAll = nil; want non-nil value")
	}
}

func TestPlan_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	p := culqi.Plan{}
	res, err := p.Update("pln_test_oFvWoKSAZOAH1weu", map[string]string{"tipo_plan": "universitario"})
	if err != nil {
		t.Fatalf("Plan.Update() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}

	if res.Metadata["tipo_plan"] != "universitario" {
		t.Errorf(`Plan.Metadata["tipo_plan"] = %s; want = %q`, res.Metadata["tipo_plan"], "universitario")
	}
}

func TestPlan_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	p := culqi.Plan{}
	err := p.Delete("pln_test_W11YcJOCx4CP1XTv")
	if err != nil {
		t.Fatalf("Plan.Delete() err = %v; want = %v", err, nil)
	}
}
