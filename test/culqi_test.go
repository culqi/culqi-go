package culqi_test

import (
	"flag"
	//"testing"
	culqi "github.com/culqi/culqi-go"
)

var (
	publicKey, secretKey string
	encryptiondData      = []byte(`{		
		"rsa_public_key": "INSERTE_SU_LLAVE_RSA",
		"rsa_id": "INSERTE_SU_RSA_ID"
	}`)
)

func init() {
	flag.StringVar(&publicKey, "public_key", "INSERTE_SU_LLAVE_PUBLICA", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
	flag.StringVar(&secretKey, "secret_key", "INSERTE_SU_LLAVE_PRIVADA", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	culqi.Key(publicKey, secretKey)
}

/*
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
		culqi.Key(tc.secretKey)
		want := culqi.GetKey()
		if tc.publicKey != want.Key ||
			tc.secretKey != want.Key {
			t.Errorf("New(pKey: %q, sKey: %q); want = %q", tc.publicKey, tc.secretKey, want)
		}
	}
}
*/
