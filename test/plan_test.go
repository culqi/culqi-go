package culqi_test

import (
	"fmt"
	"net/url"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

//go test -run TestPlan_Create test/plan_test.go

func TestPlan_Create(t *testing.T) {
	var funcName string = "TestPlan_Create"
	logStartTest(funcName)
	code, res, err := culqi.CreatePlan(getJsonPlan())

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Plan.Create() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}
	logEndTest(funcName)
}

func TestPlan_CreateEncrypt(t *testing.T) {
	var funcName string = "TestPlan_CreateEncrypt"
	logStartTest(funcName)
	code, res, err := culqi.CreatePlan(getJsonPlan(), encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Plan.CreateEncrypt() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestPlan_GetByID(t *testing.T) {
	var funcName string = "TestPlan_GetByID"
	logStartTest(funcName)

	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)

	var jsonData = []byte(``)
	code, res, err := culqi.GetByIDPlan(idPlan, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Plan.GetByID() err = %v; want = %v", err, nil)
	}
	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestPlan_GetAll(t *testing.T) {
	var funcName string = "TestPlan_GetAll"
	logStartTest(funcName)

	params := url.Values{}
	params.Set("limit", "4")
	//params.Set("before", "pln_live_qnJOtJiuGT88dAa5")
	//params.Set("after", "pln_live_c6cm1JuefM0WVkli")
	//params.Set("min_amount", "300")
	//params.Set("max_amount", "500000")
	//params.Set("status", "1")
	//params.Set("creation_date_from", "1712692203")
	//params.Set("creation_date_to", "1712692203")
	var jsonData = []byte(``)
	code, res, err := culqi.GetAllPlan(params, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Plan.GetAll() err = %v; want = %v", err, nil)
	}
	if res == "" {
		t.Fatalf("ResponsePlanAll = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestPlan_Update(t *testing.T) {
	var funcName string = "TestPlan_Update"
	logStartTest(funcName)

	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)
	code, res, err := culqi.UpdatePlan(idPlan, jsonDataUpdatePlan)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Plan.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}

	logEndTest(funcName)
}
func TestPlan_UpdateEncrypt(t *testing.T) {
	var funcName string = "TestPlan_UpdateEncrypt"
	logStartTest(funcName)

	var idPlan string
	idPlan = GetIdPlan(encryptiondData...)
	fmt.Println(idPlan)
	code, res, err := culqi.UpdatePlan(idPlan, jsonDataUpdatePlan, encryptiondData...)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Plan.Update() err = %v; want = %v", err, nil)
	}

	if res == "" {
		t.Fatalf("ResponsePlan = nil; want non-nil value")
	}

	logEndTest(funcName)
}

func TestPlan_Delete(t *testing.T) {
	var funcName string = "TestPlan_Delete"
	logStartTest(funcName)

	var idPlan string
	idPlan = GetIdPlan()
	fmt.Println(idPlan)

	var jsonData = []byte(``)
	code, res, err := culqi.DeletePlan(idPlan, jsonData)

	fmt.Println(code)
	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		t.Fatalf("Plan.Delete() err = %v; want = %v", err, nil)
	}

	logEndTest(funcName)
}
