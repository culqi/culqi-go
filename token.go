package culqi

import (
	"bytes"
	"io/ioutil"
	"net/url"
)

const (
	tokensURL = baseURLToken + "/tokens"
)

// Create método para crear un token
func CreateToken(body []byte, encryptionData []byte) (string, error) {

	res, err := do("POST", tokensURL, nil, bytes.NewBuffer(body), encryptionData)
	if err != nil {
		return "", err
	}
	response, _ := ioutil.ReadAll(bytes.NewReader(res))

	return string(response), nil
}

// CreateYape Create método para crear un token yape
func CreateYape(body []byte, encryptionData []byte) (string, error) {

	res, err := do("POST", tokensURL+"/yape", nil, bytes.NewBuffer(body), encryptionData)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un token
func UpdateToken(id string, body []byte, encryptionData []byte) (string, error) {

	res, err := do("PATCH", baseURL+"/tokens/"+id, nil, bytes.NewBuffer(body), encryptionData)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetByID método para obtener un token por id
func GetByIDToken(id string, body []byte) (string, error) {
	if id == "" {
		return " ", ErrParameter
	}

	res, err := do("GET", baseURL+"/tokens/"+id, nil, bytes.NewBuffer(body), nil)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetAll método para obtener la lista de tokens
func GetAllToken(queryParams url.Values, body []byte) (string, error) {
	res, err := do("GET", baseURL+"/tokens", queryParams, bytes.NewBuffer(body), nil)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}
