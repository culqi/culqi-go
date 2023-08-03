package culqi_test

import (
	"flag"
	//"testing"
	culqi "github.com/culqi/culqi-go"
)

var (
	publicKey, secretKey string
	encryptiondData      []byte
)

func init() {
	flag.StringVar(&publicKey, "public_key", "pk_live_889113cd74ecfc55", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
	flag.StringVar(&secretKey, "secret_key", "sk_live_34a07dcb6d4c7e39", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	culqi.Key(publicKey, secretKey)

	rsa_id := "508fc232-0a9d-4fc0-a192-364a0b782b89"
	rsa_public_key := "-----BEGIN PUBLIC KEY-----\n" +
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYp0451xITpczkBrl5Goxkh7m1oynj8eDHypIn7HmbyoNJd8cS4OsT850hIDBwYmFuwmxF1YAJS8Cd2nes7fjCHh+7oNqgNKxM2P2NLaeo4Uz6n9Lu4KKSxTiIT7BHiSryC0+Dic91XLH7ZTzrfryxigsc+ZNndv0fQLOW2i6OhwIDAQAB\n" +
		"-----END PUBLIC KEY-----"

	encryptiondData = []byte(`{		
		"rsa_public_key": "` + rsa_public_key + `",
		"rsa_id":  "` + rsa_id + `"
	}`)

	/*
		encryptiondData = []byte(`{
				"rsa_public_key": "-----BEGIN PUBLIC KEY-----
				MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDswQycch0x/7GZ0oFojkWCYv+gr5CyfBKXc3Izq+btIEMCrkDrIsz4Lnl5E3FSD7/htFn1oE84SaDKl5DgbNoev3pMC7MDDgdCFrHODOp7aXwjG8NaiCbiymyBglXyEN28hLvgHpvZmAn6KFo0lMGuKnz8HiuTfpBl6HpD6+02SQIDAQAB
		-----END PUBLIC KEY-----",
				"rsa_id": "de35e120-e297-4b96-97ef-10a43423ddec"
			}`)
	*/
}
