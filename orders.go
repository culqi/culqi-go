package culqi

import (
	"net/url"
)

const (
	ordersURL = baseURL + "/orders"
)

// Create método para crear una orden
func CreateOrder(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(ordersURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una orden por id
func GetByIDOrder(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(ordersURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de Ordenes
func GetAllOrder(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(ordersURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una orden
func UpdateOrder(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(ordersURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar una orden
func DeleteOrder(id string, body []byte) (int, string, error) {
	statusCode, res, err := Delete(ordersURL, id, body)
	return statusCode, res, err
}

// Confirm método para confirmar una orden
func ConfirmOrder(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(ordersURL+"/"+id+"/confirm", body, encryptionData...)
	return statusCode, res, err
}

// Confirm método para confirmar una orden por tipo
func ConfirmTipoOrder(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(ordersURL+"/confirm", body, encryptionData...)
	return statusCode, res, err
}
