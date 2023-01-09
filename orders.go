package culqi

import (
	"bytes"
	"net/url"
)

const (
	ordersURL = baseURL + "/orders"
)

// Create método para crear una orden
func CreateOrder(body []byte) ([]byte, error) {
	res, err := do("POST", ordersURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID método para obtener una orden por id
func GetByIDOrder(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", ordersURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAll método para obtener la lista de Ordenes
func GetAllOrder(queryParams url.Values) ([]byte, error) {
	res, err := do("GET", ordersURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una orden
func UpdateOrder(id string, body []byte) ([]byte, error) {

	res, err := do("PATCH", ordersURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete método para eliminar una orden
func DeleteOrder(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("DELETE", ordersURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Confirm método para confirmar una orden
func ConfirmOrder(id string) ([]byte, error) {
	res, err := do("POST", ordersURL+"/"+id+"/confirm", nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Confirm método para confirmar una orden por tipo
func ConfirmTipoOrder(body []byte) ([]byte, error) {
	res, err := do("POST", ordersURL+"/confirm", nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}
