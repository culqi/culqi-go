package culqi_test

import (
	"strings"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

func TestToken_Create(t *testing.T) {
	if secretKey == "" {
		t.Skip("No se indicó una llave privada")
	}

	culqi.Key(publicKey)
	var jsonData = []byte(`{		
		"card_number": "4456530000001096",
		"cvv": "111",
		"expiration_month": "07",
		"expiration_year": "2023",
		"email": "prueba@culqi.com"
	}`)
	res, err := culqi.CreateToken(jsonData)
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

	culqi.Key(publicKey)
	var jsonData = []byte(`{		
		"amount":      36200,
		"fingerprint": "86d3c875769bf62b0471b47853bfda77",
		"number_phone": "900000001",
		"otp":         "425251"
	}`)

	res, err := culqi.CreateYape(jsonData)
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
