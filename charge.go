package culqi

import (
	"bytes"
	"net/url"
)

const (
	chargesURL = baseURL + "/charges"
)

// Create método para crear un cargo
func CreateCharge(body []byte, encryptionData ...byte) (string, error) {
	res, err := do("POST", chargesURL, nil, bytes.NewBuffer(body), encryptionData...)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetByID método para obtener un cargo por id
func GetByICharge(id string, body []byte) (string, error) {
	if id == "" {
		return "", ErrParameter
	}

	res, err := do("GET", chargesURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetAll método para obtener la lista de Cargos
func GetAllCharge(queryParams url.Values, body []byte) (string, error) {
	res, err := do("GET", chargesURL, queryParams, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	response := string(res[:])

	return response, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un cargo
func UpdateCharge(id string, body []byte, encryptionData ...byte) (string, error) {
	res, err := do("PATCH", chargesURL+"/"+id, nil, bytes.NewBuffer(body), encryptionData...)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}
