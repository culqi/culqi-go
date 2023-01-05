package culqi_test

import (
	"fmt"
	"strings"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestToken_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, publicKey)
	c := culqi.Token{
		CardNumber:      "4456530000001096",
		Cvv:             "111",
		ExpirationMonth: "09",
		ExpirationYear:  "2023",
		Email:           "jordan.diaz@culqi.com",
		Metadata:        map[string]string{"coment": "Tarjeta de prueba alexis"},
	}

	res, err := c.Create()
	fmt.Println(err)
	if err != nil {
		t.Fatalf("Token.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseToken = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "tkn_test_") {
		t.Errorf("Token.ID = %s; want prefix = %q", res.ID, "tkn_test_")
	}
}

func TestToken_CreateYape(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey, secretKey)
	c := culqi.TokenYape{
		Amount:      36200,
		FingerPrint: "86d3c875769bf62b0471b47853bfda77",
		NumberPhone: "900000001",
		Otp:         "425251",
	}

	res, err := c.CreateYape()
	if err != nil {
		t.Fatalf("Token.Create() err = %v; want = %v", err, nil)
	}

	if res == nil {
		t.Fatalf("ResponseTokenYape = nil; want non-nil value")
	}

	if !strings.HasPrefix(res.ID, "ype_test_") {
		t.Errorf("Token.ID = %s; want prefix = %q", res.ID, "ype_test_")
	}
}
