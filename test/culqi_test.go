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
	flag.StringVar(&publicKey, "public_key", "INSERTE_SU_LLAVE_PUBLICA", "Su test publicKey para Culqi API. Utilizado para crear Tokens")
	flag.StringVar(&secretKey, "secret_key", "INSERTE_SU_LLAVE_PRIVADA", "Su test secretKey para Culqi API. Si esta presente, los test de integración serán ejecutados con esta llave.")
	culqi.Key(publicKey, secretKey)

	// Quitar comentario y añadir llaves RSA para flujo de pruebas RSA

	rsaID := "INSERTE_SU_RSA_ID"
	rsaPublicKey := "INSERTE_SU_LLAVE_RSA"

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
