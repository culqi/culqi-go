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
	flag.StringVar(&secretKey, "secret_key", "", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	flag.StringVar(&publicKey, "public_key", "", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
}

func TestNew(t *testing.T) {
	testCases := []struct {
		publicKey string
		secretKey string
	}{
		{"pk_test_nKjwQNbErg", "sk_test_eLc81mMel1"},
		{"pk_test_PNTiazAfPS", "sk_test_9EMXDrQkqX"},
		{"pk_test_Ht1Vbvyy4x", "sk_test_swr6XRR5yH"},
		{"pk_test_EMwRgHe3x4", "sk_test_8PzpFp1lpf"},
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
