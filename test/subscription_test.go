package culqi_test

import (
	"net/url"
	"strings"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestSubscription_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)
	s := culqi.Subscription{
		CardID:   "crd_test_Qe0HG7VTfmTdFvgr",
		PlanID:   "pln_test_oFvWoKSAZOAH1weu",
		Metadata: map[string]string{"user_id": "723"},
	}

	res, err := s.Create()
	if err != nil {
		t.Fatalf("Subscription.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "sub_test_") {
		t.Errorf("Subscription.ID = %s; want prefix = %q", res.ID, "sub_test_")
	}
}

func TestSubscription_GetByID(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	c := culqi.Subscription{}
	res, err := c.GetByID("sub_test_QjfgsHMPghROVZXa")
	if err != nil {
		t.Fatalf("Subscription.GetByID() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}
}

func TestSubscription_GetAll(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	c := culqi.Subscription{}
	params := url.Values{}
	params.Set("limit", "4")

	res, err := c.GetAll(params)
	if err != nil {
		t.Fatalf("Subscription.GetAll() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseSubscriptionAll = nil; want non-nil value")
	}
}

func TestSubscription_Update(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	c := culqi.Subscription{}
	res, err := c.Update("sub_test_QjfgsHMPghROVZXa", map[string]string{"pais": "Colombia"})
	if err != nil {
		t.Fatalf("Subscription.Update() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseSubscription = nil; want non-nil value")
	}

	if res.Metadata["pais"] != "Colombia" {
		t.Errorf(`Subscription.Metadata["pais"] = %s; want = %q`, res.Metadata["pais"], "Colombia")
	}
}

func TestSubscription_Delete(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)

	p := culqi.Subscription{}
	err := p.Delete("sub_test_eY5TIGm70OMiCW89")
	if err != nil {
		t.Fatalf("Subscription.Delete() err = %v; want = %v", err, nil)
	}
}
