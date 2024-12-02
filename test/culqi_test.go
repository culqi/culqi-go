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

	rsaID := "2ce8e080-4cdf-454f-9c1b-b0b7b98b73d7"
	rsaPublicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC6ncJ+6Lzaomc3bA7fmk7bZO5b
CkuMthSvWYpOL9/Y5qs/DhE1bp1AlA0QXAJztcrhiCk+giUNdPgmT3oT1yfAAWW8
ahvol8QjhJJ1FAxNJjcIErN3EOeoP8F/2U1ESJkVXArcwq2LnFnmHtwYcCqU4ZQq
J7mQU6lQ6ezMVuKAeQIDAQAB
-----END PUBLIC KEY-----`

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
