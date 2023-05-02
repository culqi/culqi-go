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
	flag.StringVar(&secretKey, "secret_key", "sk_live_34a07dcb6d4c7e39", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	flag.StringVar(&publicKey, "public_key", "pk_live_889113cd74ecfc55", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
}

func TestNew(t *testing.T) {
	testCases := []struct {
		publicKey string
		secretKey string
	}{
		{"pk_live_889113cd74ecfc55", "sk_live_34a07dcb6d4c7e39"},
		{"pk_live_889113cd74ecfc55", "sk_live_34a07dcb6d4c7e39"},
		{"pk_live_889113cd74ecfc55", "sk_live_34a07dcb6d4c7e39"},
		{"pk_live_889113cd74ecfc55", "sk_live_34a07dcb6d4c7e39"},
	}

	for _, tc := range testCases {
		culqi.Key(tc.secretKey)
		want := culqi.GetKey()
		if tc.publicKey != want.Key ||
			tc.secretKey != want.Key {
			t.Errorf("New(pKey: %q, sKey: %q); want = %q", tc.publicKey, tc.secretKey, want)
		}
	}
}
