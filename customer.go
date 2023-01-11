package culqi

import (
	"bytes"
	"net/url"
)

const (
	customerURL = baseURL + "/customers"
)

// Create método para crear un cliente
func CreateCustomer(body []byte) (string, error) {
	res, err := do("POST", customerURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetByID método para obtener un cliente por id
func GetByIDCustomer(id string) (string, error) {
	if id == "" {
		return "", ErrParameter
	}

	res, err := do("GET", customerURL+"/"+id, nil, nil)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetAll método para obtener la lista de clientes
func GetAllCustomer(queryParams url.Values) (string, error) {
	res, err := do("GET", customerURL, queryParams, nil)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un cliente
func UpdateCustomer(id string, body []byte) (string, error) {
	res, err := do("PATCH", customerURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil

}

// Delete método para eliminar un cliente por id
func DeleteCustomer(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", customerURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
