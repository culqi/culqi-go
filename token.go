package culqi

import (
	"encoding/json"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	tokensSecureURL = baseURLSecure + "/tokens"
	tokensApiURL    = baseURL + "/tokens"
)

// Create método para crear un token
func CreateToken(body []byte, encryptionData ...byte) (int, string, error) {
	// Unmarshal the body into a map
	var data map[string]string
	err := json.Unmarshal(body, &data)
	if err != nil {
		return 0, "", err
	}

	// Perform validation
	validator := utils.NewTokenValidation()
	err = validator.Create(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}

	statusCode, res, err := Create(tokensSecureURL, body, encryptionData...)
	return statusCode, res, err
}

// CreateYape Create método para crear un token yape
func CreateYape(body []byte, encryptionData ...byte) (int, string, error) {
	var data map[string]interface{}
	errassign := json.Unmarshal(body, &data)
	if errassign != nil {
		return ErrorBadRequest, "", errassign
	}
	err := utils.CreateTokenYape(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := Create(tokensSecureURL+"/yape", body, encryptionData...)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un token
func UpdateToken(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "tkn")
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := Update(tokensApiURL, id, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener un token por id
func GetByIDToken(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "tkn")
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := GetById(tokensApiURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de tokens
func GetAllToken(queryParams url.Values, body []byte) (int, string, error) {
	data := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			data[key] = values[0] // Taking only the first value for each key.
		}
	}

	err := utils.TokenListValidation(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}

	statusCode, res, err := GetAll(tokensApiURL, queryParams, body)
	return statusCode, res, err
}
