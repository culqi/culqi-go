package culqi_test

import (
	"flag"
	"testing"

	culqi "github.com/culqi/culqi-go"
)

var (
	publicKey, secretKey, secretKeyToken string
)

func init() {
	flag.StringVar(&secretKeyToken, "secret_keyToken", "INSERTE_SU_LLAVE_PUBLICA", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	flag.StringVar(&secretKey, "secret_key", "INSERTE_SU_LLAVE_PRIVADA", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	flag.StringVar(&publicKey, "public_key", "INSERTE_SU_LLAVE_PUBLICA", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
}

func TestNew(t *testing.T) {
	testCases := []struct {
		publicKey string
		secretKey string
	}{
		{"INSERTE_SU_LLAVE_PUBLICA", "INSERTE_SU_LLAVE_PRIVADA"},
		{"INSERTE_SU_LLAVE_PUBLICA", "INSERTE_SU_LLAVE_PRIVADA"},
		{"INSERTE_SU_LLAVE_PUBLICA", "INSERTE_SU_LLAVE_PRIVADA"},
		{"INSERTE_SU_LLAVE_PUBLICA", "INSERTE_SU_LLAVE_PRIVADA"},
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
