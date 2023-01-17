package culqi

import (
	"bytes"
	"net/url"
)

const (
	ordersURL = baseURL + "/orders"
)

// Create método para crear una orden
func CreateOrder(body []byte) (string, error) {
	res, err := do("POST", ordersURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetByID método para obtener una orden por id
func GetByIDOrder(id string, body []byte) (string, error) {
	if id == "" {
		return "", ErrParameter
	}

	res, err := do("GET", ordersURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetAll método para obtener la lista de Ordenes
func GetAllOrder(queryParams url.Values, body []byte) (string, error) {
	res, err := do("GET", ordersURL, queryParams, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una orden
func UpdateOrder(id string, body []byte) (string, error) {

	res, err := do("PATCH", ordersURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Delete método para eliminar una orden
func DeleteOrder(id string, body []byte) (string, error) {
	if id == "" {
		return "", ErrParameter
	}

	res, err := do("DELETE", ordersURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Confirm método para confirmar una orden
func ConfirmOrder(id string, body []byte) (string, error) {
	res, err := do("POST", ordersURL+"/"+id+"/confirm", nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Confirm método para confirmar una orden por tipo
func ConfirmTipoOrder(body []byte) (string, error) {
	res, err := do("POST", ordersURL+"/confirm", nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}
