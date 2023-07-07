package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestPlan_Create(t *testing.T) {
	_, res, err := culqi.CreatePlan(jsonDataPlan)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Plan.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
}

func TestPlan_GetByID(t *testing.T) {

	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)

	var jsonData = []byte(``)
	_, res, err := culqi.GetByIDPlan(idPlan, jsonData)
	if err != nil {
		t.Fatalf("Plan.GetByID() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
}

func TestPlan_GetAll(t *testing.T) {

	params := url.Values{}
	params.Set("limit", "4")
	var jsonData = []byte(``)
	_, res, err := culqi.GetAllPlan(params, jsonData)
	if err != nil {
		t.Fatalf("Plan.GetAll() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlanAll = nil; want non-nil value")
	}
}

func TestPlan_Update(t *testing.T) {
	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)

	var jsonData = []byte(`{
		"metadata": {
		"descripcion": "Este es un plan simple."
		}
	}`)

	_, res, err := culqi.UpdatePlan(idPlan, jsonData)
	if err != nil {
		t.Fatalf("Plan.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
}

func TestPlan_Delete(t *testing.T) {
	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)

	var jsonData = []byte(``)
	_, res, err := culqi.DeletePlan(idPlan, jsonData)
	fmt.Println(res)
	if err != nil {
		t.Fatalf("Plan.Delete() err = %v; want = %v", err, nil)
	}
}
