package culqi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	utils "github.com/culqi/culqi-go/utils"
	culqi "github.com/culqi/culqi-go/utils/encryption/rsa_aes"
)

const (
	apiVersion    = "v2.0"
	baseURL       = "https://api.culqi.com/v2"
	baseURLSecure = "https://secure.culqi.com/v2"
)

// Errors API
var (
	ErrInvalidRequest = errors.New("La petición tiene una sintaxis inválida")
	ErrAuthentication = errors.New("La petición no pudo ser procesada debido a problemas con las llaves")
	ErrParameter      = errors.New("Algún parámetro de la petición es inválido")
	ErrCard           = errors.New("No se pudo realizar el cargo a una tarjeta")
	ErrLimitAPI       = errors.New("Estás haciendo muchas peticiones rápidamente al API o superaste tu límite designado")
	ErrResource       = errors.New("El recurso no puede ser encontrado, es inválido o tiene un estado diferente al permitido")
	ErrAPI            = errors.New("Error interno del servidor de Culqi")
	ErrUnexpected     = errors.New("Error inesperado, el código de respuesta no se encuentra controlado")
	ErrKey            = errors.New("El formato de llaves debe iniciar con pk_test, pk_live, sk_test o sk_live")
	ErrorGenerico     = 502
	ErrorBadRequest   = 400
)

// WrapperResponse respuesta generica para respuestas GetAll
type WrapperResponse struct {
	Paging struct {
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Cursors  struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

// create
func do(method, endpoint string, params url.Values, body io.Reader, encryptionData ...byte) (int, []byte, error) {
	errKey := CheckKey(keyInstance.publicKey, keyInstance.secretKey)
	if errKey != nil {
		return ErrorGenerico, nil, errKey
	}
	idRsaHeader := ""
	key := ""

	// Se crea un objeto con los headers a añadir
	headers := make(map[string]interface{})

	if encryptionData != nil {
		// Deserializar el JSON en un mapa
		var source map[string]interface{}

		err := json.Unmarshal(encryptionData, &source)
		if err != nil {
			log.Fatalf("Error al deserializar el JSON: %v", err)
		}

		// Crear un objeto vacío
		destination := make(map[string]interface{})

		// Iterar sobre las claves del mapa original
		for key, value := range source {
			switch key {
			case "custom_headers":
				fmt.Println(value)
				for header, HeaderVale := range value.(map[string]interface{}) {
					fmt.Println(`Adding header '`, header, `' with value '`, HeaderVale, `'`)
					headers[header] = HeaderVale.(string)
				}
			case "rsa_public_key":
				destination["rsa_public_key"] = value
			case "rsa_id":
				destination["rsa_id"] = value
			default:
				// Manejo para claves desconocidas (si es necesario)
				fmt.Printf("Clave '%s' no está manejada en el switch\n", key)
			}
		}

		// Convertir el objeto vacío actualizado de nuevo a JSON
		updatedJSON, err := json.Marshal(destination)
		if err != nil {
			log.Fatalf("Error al serializar el JSON: %v", err)
		}

		encryptionData = updatedJSON

		if string(updatedJSON) != "{}" {
			body, idRsaHeader = culqi.Encrypt(body, updatedJSON)
		}
	} else {
		fmt.Println("No encryption data")
	}

	if len(params) != 0 {
		endpoint += "?" + params.Encode()
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	fmt.Println(endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(buf.Bytes()))
	if err != nil {
		return ErrorGenerico, nil, err
	}

	for header, headerValue := range headers {
		fmt.Println(`Adding header '`, header, `' with value '`, headerValue, `'`)
		req.Header.Set(header, headerValue.(string))
	}

	if method == "POST" {
		if strings.Contains(endpoint, "v2/tokens") || strings.Contains(endpoint, "confirm") {
			key = keyInstance.publicKey
		} else {
			key = keyInstance.secretKey
		}
	} else {
		key = keyInstance.secretKey
	}

	env := utils.XCulqiEnvLive

	if strings.Contains(key, "test") {
		env = utils.XCulqiEnvTest
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("x-culqi-env", env)
	req.Header.Set("x-api-version", utils.XApiVersion)
	req.Header.Set("x-culqi-client", utils.XCulqiClient)
	req.Header.Set("x-culqi-client-version", utils.XCulqiClientVersion)
	if idRsaHeader != "" {
		req.Header.Set("x-culqi-rsa-id", idRsaHeader)
	}

	c := &http.Client{}

	res, err := c.Do(req)
	if err != nil {
		return ErrorGenerico, nil, err
	}
	defer res.Body.Close()

	obj, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ErrorGenerico, nil, err
	}
	
	fmt.Println(res.StatusCode)
	switch res.StatusCode {
	case 400:
		err = ErrInvalidRequest
	case 401:
		err = ErrAuthentication
	case 422:
		err = ErrParameter
	case 402:
		err = ErrCard
	case 429:
		err = ErrLimitAPI
	case 404:
		err = ErrResource
	case 500, 503:
		err = ErrAPI
	}

	if err != nil {
		err = fmt.Errorf("%v: %s", err, string(obj))
		return res.StatusCode, nil, err
	}

	if res.StatusCode >= 200 && res.StatusCode <= 206 {
		return res.StatusCode, obj, nil
	}

	return ErrorGenerico, nil, ErrUnexpected
}

//Funciones genericas

func Create(URL string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := do("POST", URL, nil, bytes.NewBuffer(body), encryptionData...)
	if err != nil {
		return statusCode, "", err
	}
	response := string(res[:])
	fmt.Println(response)
	return statusCode, response, nil
}

func GetById(URL string, id string, body []byte) (int, string, error) {
	if id == "" {
		return ErrorGenerico, "", ErrParameter
	}

	statusCode, res, err := do("GET", URL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return statusCode, "", err
	}
	response := string(res[:])
	fmt.Println(response)
	return statusCode, response, nil
}

func GetAll(URL string, queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := do("GET", URL, queryParams, bytes.NewBuffer(body))
	if err != nil {
		return statusCode, "", err
	}

	response := string(res[:])

	return statusCode, response, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un cargo
func Update(URL string, id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := do("PATCH", URL+"/"+id, nil, bytes.NewBuffer(body), encryptionData...)
	fmt.Println(URL + "/" + id)
	if err != nil {
		return statusCode, "", err
	}
	response := string(res[:])
	fmt.Println(response)
	return statusCode, response, nil
}

func Delete(URL string, id string, body []byte) (int, string, error) {
	if id == "" {
		return ErrorGenerico, "", ErrParameter
	}

	statusCode, res, err := do("DELETE", URL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return statusCode, "", err
	}
	response := string(res[:])

	return statusCode, response, nil
}

func CheckKey(publicKey string, secretKey string) error {
	if !strings.HasPrefix(publicKey, "pk_test_") &&
		!strings.HasPrefix(publicKey, "pk_live_") {
		return ErrKey
	}

	if !strings.HasPrefix(secretKey, "sk_test_") &&
		!strings.HasPrefix(secretKey, "sk_live_") {
		return ErrKey
	}

	return nil
}
