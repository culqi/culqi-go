package culqi

import (
	"bytes"
	"net/url"
)

const (
	cardURL = baseURL + "/cards"
)

// Create método para crear una tarjeta
func CreateCard(body []byte) (string, error) {

	res, err := do("POST", cardURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetByID método para obtener una tarjeta por id
func GetByIDCard(id string, body []byte) (string, error) {
	if id == "" {
		return "", ErrParameter
	}

	res, err := do("GET", cardURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetAll método para obtener la lista de las tarjetas
func GetAllCard(queryParams url.Values, body []byte) (string, error) {
	res, err := do("GET", cardURL, queryParams, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una tarjeta
func UpdateCard(id string, body []byte) (string, error) {

	res, err := do("PATCH", cardURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Delete método para eliminar una tarjeta por id
func DeleteCard(id string, body []byte) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", cardURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return nil
}
