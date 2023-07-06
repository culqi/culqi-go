package culqi

import (
	"net/url"
)

const (
	customerURL = baseURL + "/customers"
)

// Create método para crear un cliente
func CreateCustomer(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(customerURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener un cliente por id
func GetByIDCustomer(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(customerURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de clientes
func GetAllCustomer(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(customerURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un cliente
func UpdateCustomer(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(customerURL, id, body, encryptionData...)
	return statusCode, res, err

}

// Delete método para eliminar un cliente por id
func DeleteCustomer(id string, body []byte) (int, string, error) {
	statusCode, res, err := Delete(customerURL, id, body)
	return statusCode, res, err
}
