package culqi_test

import (
	"flag"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

var (
	publicKey, secretKey string
)

func init() {
	flag.StringVar(&secretKey, "secret_key", "pk_test_90667d0a57d45c48", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	flag.StringVar(&publicKey, "public_key", "pk_test_90667d0a57d45c48", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
}

func TestNew(t *testing.T) {
	testCases := []struct {
		publicKey string
		secretKey string
	}{
		{"pk_test_387cc0e60fa9f7d4", "sk_test_ff27818fc60ff66a"},
		{"pk_test_387cc0e60fa9f7d4", "sk_test_ff27818fc60ff66a"},
		{"pk_test_387cc0e60fa9f7d4", "sk_test_ff27818fc60ff66a"},
		{"pk_test_387cc0e60fa9f7d4", "sk_test_ff27818fc60ff66a"},
	}

	for _, tc := range testCases {
		culqi.Key(tc.publicKey, tc.secretKey)
		want := culqi.GetKey()
		if tc.publicKey != want.PublicKey ||
			tc.secretKey != want.SecretKey {
			t.Errorf("New(pKey: %q, sKey: %q); want = %q", tc.publicKey, tc.secretKey, want)
		}
	}
}
