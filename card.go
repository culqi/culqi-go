package culqi

import (
	"bytes"
	"net/url"
)

const (
	cardURL = baseURL + "/cards"
)

// Create método para crear una tarjeta
func CreateCard(body []byte) ([]byte, error) {

	res, err := do("POST", cardURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID método para obtener una tarjeta por id
func GetByIDCard(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", cardURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAll método para obtener la lista de las tarjetas
func GetAllCard(queryParams url.Values) ([]byte, error) {
	res, err := do("GET", cardURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una tarjeta
func UpdateCard(id string, body []byte) ([]byte, error) {

	res, err := do("PATCH", cardURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete método para eliminar una tarjeta por id
func DeleteCard(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", cardURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
