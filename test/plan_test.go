package culqi_test

import (
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestPlan_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	var jsonData = []byte(`{
	  "name": "Prueba Webhook",
	  "amount": 300,
	  "currency_code": "PEN",
	  "interval": "dias",
	  "interval_count": 1,
	  "limit": 3,
	  "trial_days": 1
	}`)

	res, err := culqi.CreatePlan(jsonData)
	if err != nil {
		t.Fatalf("Plan.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
}

func TestPlan_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	res, err := culqi.GetByIDPlan("pln_test_oFvWoKSAZOAH1weu")
	if err != nil {
		t.Fatalf("Plan.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
}

func TestPlan_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	params := url.Values{}
	params.Set("limit", "4")

	res, err := culqi.GetAllPlan(params)
	if err != nil {
		t.Fatalf("Plan.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlanAll = nil; want non-nil value")
	}
}

func TestPlan_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}
	var jsonData = []byte(`{
		"metadata": {
		"descripcion": "Este es un plan simple."
		}
	}`)
	culqi.Key(secretKey)
	res, err := culqi.UpdatePlan("pln_test_oFvWoKSAZOAH1weu", jsonData)
	if err != nil {
		t.Fatalf("Plan.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
}

func TestPlan_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(secretKey)
	err := culqi.DeletePlan("pln_test_W11YcJOCx4CP1XTv")
	if err != nil {
		t.Fatalf("Plan.Delete() err = %v; want = %v", err, nil)
	}
}
