package culqi

import (
	"bytes"
	"net/url"
)

const (
	tokensURL = baseURLToken + "/tokens"
)

// Create método para crear un token
func CreateToken(body []byte) ([]byte, error) {

	res, err := do("POST", tokensURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	//myString := string(res[:])

	return res, nil
}

// Create método para crear un token yape
func CreateYape(body []byte) ([]byte, error) {

	res, err := do("POST", tokensURL+"/yape", nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un token
func UpdateToken(id string, body []byte) ([]byte, error) {

	res, err := do("PATCH", tokensURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID método para obtener un token por id
func GetByIDToken(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", tokensURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAll método para obtener la lista de tokens
func GetAllToken(queryParams url.Values) ([]byte, error) {
	res, err := do("GET", tokensURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
