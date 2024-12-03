package culqi_test

import (
	"encoding/json"
	"flag"
	"fmt"
	culqi "github.com/culqi/culqi-go"
)

var (
	publicKey, secretKey string
	encryptiondData      []byte
)

func init() {
	flag.StringVar(&publicKey, "public_key", "pk_test_e94078b9b248675d", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
	flag.StringVar(&secretKey, "secret_key", "sk_test_c2267b5b262745f0", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	culqi.Key(publicKey, secretKey)

	// Quitar comentario y añadir llaves RSA para flujo de pruebas RSA

	rsaID := "de35e120-e297-4b96-97ef-10a43423ddec"
	rsaPublicKey := "-----BEGIN PUBLIC KEY-----\n" +
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDswQycch0x/7GZ0oFojkWCYv+gr5CyfBKXc3Izq+btIEMCrkDrIsz4Lnl5E3FSD7/htFn1oE84SaDKl5DgbNoev3pMC7MDDgdCFrHODOp7aXwjG8NaiCbiymyBglXyEN28hLvgHpvZmAn6KFo0lMGuKnz8HiuTfpBl6HpD6+02SQIDAQAB\n" +
		"-----END PUBLIC KEY-----"

	// Crear el mapa para los datos
	data := map[string]string{
		"rsa_public_key": rsaPublicKey,
		"rsa_id":         rsaID,
	}

	var err error
	encryptiondData, err = json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}
}
