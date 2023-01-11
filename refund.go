package culqi

import (
	"bytes"
	"net/url"
)

const (
	refundURL = baseURL + "/refunds"
)

// Create método para crear una devolucion
func CreateRefund(body []byte) ([]byte, error) {
	res, err := do("POST", chargesURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID método para obtener una devolucion por id
func GetByIDRefund(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", refundURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAll método para obtener la lista de devoluciones
func GetAllRefund(queryParams url.Values) ([]byte, error) {
	res, err := do("GET", refundURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una devolucion
func UpdateRefund(id string, body []byte) ([]byte, error) {
	res, err := do("PATCH", chargesURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}
