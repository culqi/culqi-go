package culqi_test

import (
	"flag"
	//"testing"
	culqi "github.com/culqi/culqi-go"
)

var (
	publicKey, secretKey string
	encryptiondData      = []byte(`{		
		"rsa_public_key": "-----BEGIN PUBLIC KEY-----
		MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDswQycch0x/7GZ0oFojkWCYv+gr5CyfBKXc3Izq+btIEMCrkDrIsz4Lnl5E3FSD7/htFn1oE84SaDKl5DgbNoev3pMC7MDDgdCFrHODOp7aXwjG8NaiCbiymyBglXyEN28hLvgHpvZmAn6KFo0lMGuKnz8HiuTfpBl6HpD6+02SQIDAQAB
	-----END PUBLIC KEY-----",
		"rsa_id": "de35e120-e297-4b96-97ef-10a43423ddec"
	}`)
)

func init() {
	flag.StringVar(&publicKey, "public_key", "pk_test_e94078b9b248675d", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
	flag.StringVar(&secretKey, "secret_key", "sk_test_c2267b5b262745f0", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	culqi.Key(publicKey, secretKey)
}

/*
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
*/
