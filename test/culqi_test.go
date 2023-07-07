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
	flag.StringVar(&publicKey, "public_key", "INSERTE_SU_LLAVE_PUBLICA", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
	flag.StringVar(&secretKey, "secret_key", "INSERTE_SU_LLAVE_PRIVADA", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	culqi.Key(publicKey, secretKey)

	rsa_id := "INSERTE_SU_RSA_ID"
	rsa_public_key := "INSERTE_SU_LLAVE_RSA"

	encryptiondData = []byte(`{		
		"rsa_public_key": "` + rsa_public_key + `",
		"rsa_id":  "` + rsa_id + `"
	}`)

	/*
		encryptiondData = []byte(`{
				"rsa_public_key": "INSERTE_SU_LLAVE_RSA",
				"rsa_id": "INSERTE_SU_RSA_ID"
			}`)
	*/
}
