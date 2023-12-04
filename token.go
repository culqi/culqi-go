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
	validator := utils.NewCulqiValidation()
	err = validator.CreateTokenValidation(data)
	if err != nil {
		return 0, "", err
	}

	statusCode, res, err := Create(tokensSecureURL, body, encryptionData...)
	return statusCode, res, err
}

// CreateYape Create método para crear un token yape
func CreateYape(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(tokensSecureURL+"/yape", body, encryptionData...)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un token
func UpdateToken(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(tokensApiURL, id, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener un token por id
func GetByIDToken(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(tokensApiURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de tokens
func GetAllToken(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(tokensApiURL, queryParams, body)
	return statusCode, res, err
}
