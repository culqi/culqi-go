package culqi

import (
	"bytes"
	"net/url"
)

const (
	chargesURL = baseURL + "/charges"
)

// Create método para crear un cargo
func CreateCharge(body []byte) ([]byte, error) {
	res, err := do("POST", chargesURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID método para obtener un cargo por id
func GetByICharge(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", chargesURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAll método para obtener la lista de Cargos
func GetAllCharge(queryParams url.Values) ([]byte, error) {
	res, err := do("GET", chargesURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un cargo
func UpdateCharge(id string, body []byte) ([]byte, error) {
	res, err := do("PATCH", chargesURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}
